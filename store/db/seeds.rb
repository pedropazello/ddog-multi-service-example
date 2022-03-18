# This file should contain all the record creation needed to seed the database with its default values.
# The data can then be loaded with the bin/rails db:seed command (or created alongside the database with db:setup).
#
# Examples:
#
#   movies = Movie.create([{ name: "Star Wars" }, { name: "Lord of the Rings" }])
#   Character.create(name: "Luke", movie: movies.first)

order1 = Order.create!

2.times.each do |i|
  Product.create!(order: order1, name: "Product #{i}", price_cents: (i + 1) * 100)
end

order2 = Order.create!

2.times.each do |i|
  Product.create!(order: order2, name: "error", price_cents: (i + 1) * 100)
end
