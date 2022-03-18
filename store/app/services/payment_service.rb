class PaymentService
  attr_reader :order

  def initialize(order)
    @order = order
  end

  def execute
    Datadog.tracer.trace('PaymentService.Execute') do |span|
      span.set_tag("payload", body)

      begin
        response = RestClient.post('http://checkout_service:8080/checkouts', body)
      rescue RestClient::BadRequest => e
        span.set_tag("error", "failed to create checkout: #{e.message}")
        span.finish(e)
      end

      span.finish
    end
  end

  private

  def body
    {
      "products" => order.products.map do |product|
        {
          "name" => product.name,
          "price_cents" => product.price_cents
        }
      end
    }.to_json
  end
end
