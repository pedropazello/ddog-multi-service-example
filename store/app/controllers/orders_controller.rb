class OrdersController < ApplicationController

  skip_before_action :verify_authenticity_token

  def pay
    @order = Order.find(params[:order_id])

    PaymentService.new(@order).execute

    render json: { processing: true }
  end
end
