Rails.application.routes.draw do
  resources :orders, only: [] do
    patch :pay
  end
end
