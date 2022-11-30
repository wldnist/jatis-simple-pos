-- a. List of customers located in Irvine city.
SELECT * FROM customers WHERE city = 'Irvine';

-- b. List of customers whose order is handled by an employee named Adam Barr.
SELECT c.* 
FROM customers c
RIGHT JOIN orders o ON c.customer_id = o.customer_id
LEFT JOIN employees e ON o.employee_id = e.employee_id
WHERE CONCAT(e.first_name, " ", e.last_name) = 'Adam Barr'
GROUP BY customer_id;

-- c. List of products which are ordered by "Contonso, Ltd" Company.
SELECT p.* 
FROM products p
RIGHT JOIN order_details od ON p.product_id = od.product_id
LEFT JOIN orders o ON od.order_id = o.order_id
LEFT JOIN customers c ON o.customer_id = c.customer_id
WHERE c.company_name= 'Contonso, Ltd'
GROUP BY product_id;

-- d. List of transactions (orders) which has "UPS Ground" as shipping method.
SELECT o.* 
FROM orders o
LEFT JOIN shipping_methods sm ON o.shipping_method_id = sm.shipping_method_id
WHERE sm.shipping_method = 'UPS Ground';

-- e. List of total cost (including tax and freight charge) for every order sorted by ship date.
SELECT o.*, (SUM((od.unit_price * od.quantity) - od.discount) + o.freight_charge + o.taxes) AS total_cost
FROM orders o
JOIN order_details od ON o.order_id = od.order_id
GROUP BY od.order_id
ORDER BY o.ship_date;