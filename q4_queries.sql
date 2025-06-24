/*
	EPBC Assessment Q4 SQL Submission
	Date: 2025-06-23
	Name: Mac Cooper
*/

 -- Q4.1:	Best-selling product for a given time period, based on total sales(i.e. It has generated the most revenue), and unit sales(i.e it has sold the most units)?

DECLARE @startDate DATETIME = '2025-06-23';
DECLARE @endDate DATETIME = '2025-06-01';

SELECT TOP 1
	li.ProductID,
	SUM(li.Quantity) as GrossUnitsSold,
	SUM(li.LineItemUnitPrice * li.Qunatity) AS TotalRevenue
FROM LineItem AS li
JOIN [Order] AS o ON li.OrderID = o.OrderID
WHERE o.OrderCreatedDate BETWEEN @startDate AND @endDate 
GROUP BY li.ProductID
ORDER BY TotalRevenue DESC

 -- Q4.2:	Who are the top 5 customers based on their total net sales (i.e. net sales would be the total amount of sales they have placed minus returns they have made)	

SELECT TOP 5
	c.CustomerID,
	c.FirstName,
	c.LastName,
	-- (Total Sales - Total returns)
	(
		SELECT 
			SUM(o.OrderTotal)
		FROM [Order] AS o
		WHERE o.CustomerID = c.CustomerID

	) - 
	(
		SELECT
			SUM(ri.AmountReturend)
		FROM [ReturnItem] AS ri
		JOIN [Order] AS ON ri.OrderID = o.OrderID
		WHERE o.CustomerID = c.CustomerID
	) AS NetSales
FROM Customer AS c
ORDER BY NetSales DESC


/* Q4.3:	Categorizing customers into the age brackets [(0,17),(18,29),(30,45),(46-65),(65+)], how would we get the following:
		Q4.3.1: * Which age group generated the most sales
		Q4.3.2: * What was the top country from sales perspective for each age bracket and what was the total sales?
*/

WITH CustomerAgeSales AS (
	SELECT
		c.CustomerID,
		o.OrderTotal,
			CASE
				WHEN DATEDIFF(YEAR, c.DOB, GETDATE()) < 18 THEN 'Under 18'
				WHEN DATEDIFF(YEAR, c.DOB, GETDATE()) BETWEEN 18 AND 29 THEN '18-29'
				WHEN DATEDIFF(YEAR, c.DOB, GETDATE()) BETWEEN 30 AND 45 THEN '30-45'
				WHEN DATEDIFF(YEAR, c.DOB, GETDATE()) BETWEEN 46 AND 65 THEN '46-65'
				ELSE '65+'
			END AS AgeBracket,
		FROM Customer AS c
		JOIN [Order] AS o ON c.CustomerID = o.CustomerID
)
SELECT TOP 1
	AgeBracket,
	SUM(OrderTotal) AS TotalSales
FROM CustomerAgeSales
GROUP BY AgeBracket
ORDER BY TotalSales DESC;




/* Q4.4:	Provide the data for a histogram that shows the number of orders by number of days it took to fulfill the order after it was placed. Ex:
		* < 1 day: 50 orders
		* 1-2 days: 350 orders
		* 2-3 days: 400 orders
		* 3-4 days: 200 orders
		* > 4 days: 75 orders
*/

WITH FulfillDelays AS (
	SELECT
		DATEDIFF(day, OrderCreatedDate, OrderFullfilledDate) AS FulfillmentDays
	FROM [Order]
	WHERE OrderFulfilledDate IS NOT NULL

)

SELECT 
	CASE
		WHEN FulfillmentDays < 1 THEN 'Less than a day'
		WHEN FulfillmentDays BETWEEN 1 AND 2 THEN '1-2 days'
		WHEN FulfillmentDays BETWEEN 2 AND 3 THEN '2-3 days'
		WHEN FulfillmentDays BETWEEN 3 AND 4 THEN '3-4 days'
		ELSE 'Greater Than 4 days'
	END AS FulfillmentBuckets,
	COUNT(*) AS NumberOfOrders
FROM FulfillmentDelays
GROUP BY FulfillmentBuckets
ORDER BY FulfillmentDelays


 -- Q4.5:	What % of items that were ordered included a discount from the full price?

SELECT CAST(100.0 * 13 / 37 AS DECIMAL(5,2)) AS Result;
