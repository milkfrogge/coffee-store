SELECT product.id, product.name, product.description, product.counter, product.price, product.created_at, category.id, category.name, su.arr
from product
         join category on product.category = category.id
         left join (select array_agg(url) as arr, product_images.product_id as id
                    from product_images
                             join product on product_images.product_id = product.id
                    group by product_images.product_id) as su on product.id = su.id
where product.id = '00000000-0000-0000-0000-000000000002'