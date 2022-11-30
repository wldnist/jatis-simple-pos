SELECT od.order_id, 
    od.product_id, 
    p.product_name, 
    SUM(od.quantity) as quantity, 
    SUM(od.unit_price) AS unit_price, 
    SUM(od.discount) AS discount, 
    CONCAT(c.first_name, " ", c.last_name) AS customer_name, 
    CONCAT(e.first_name, " ", e.last_name) AS employee_name, 
    sm.shipping_method, 
    SUM(((od.quantity * od.unit_price) - od.discount)) AS sub_total
FROM order_details od
LEFT JOIN products p ON od.product_id = p.product_id
LEFT JOIN orders o ON od.order_id = o.order_id
LEFT JOIN customers c ON o.customer_id = c.customer_id
LEFT JOIN employees e ON o.employee_id = e.employee_id
LEFT JOIN shipping_methods sm ON o.shipping_method_id = sm.shipping_method_id
GROUP BY od.order_id, od.product_id;