class CheckoutCompleted
  include Shoryuken::Worker

  QUEUE = 'http://localstack:4566/000000000000/checkout-completed-queue'

  shoryuken_options queue: QUEUE, auto_delete: true

  def perform(sqs_msg, msg_body)
    puts "CheckoutCompleted: #{msg_body}"
    puts "TraceCarrier: #{sqs_msg.message_attributes.dig('TraceCarrier', 'string_value')}"

    carrier = JSON.parse(sqs_msg.message_attributes.dig('TraceCarrier', 'string_value'))

    context = ::Datadog::Context.new(
      trace_id: carrier["x-datadog-trace-id"].to_i(10),
      span_id: carrier["x-datadog-parent-id"].to_i(10),
      sampled: carrier["x-datadog-sampling-priority"] == "1",
    )

    Datadog.tracer.trace('CheckoutCompleted.perform', { child_of: context, service: 'store-shoryuken' }) do |span|
      span.set_tag('msg_body', msg_body)
      span.set_tag('queue', QUEUE)
    end
  end
end
