require 'active_record'
require 'rest_client'
require 'ddtrace'

Datadog.configure do |c|
  c.use :active_record, service_name: 'store-postgres'
  c.use :rest_client, service_name: 'store-http-request'
  c.use :shoryuken, service_name: 'store-shoryuken'
end
