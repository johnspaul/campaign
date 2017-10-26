
CREATE TABLE "campaign"
(
  campaignId VARCHAR(50) NOT NULL PRIMARY KEY,
  productId VARCHAR(50),
  startDate DATE ,
  endDate DATE
);
CREATE  TABLE "products"
(
  productId VARCHAR(50) NOT NULL PRIMARY KEY ,
  productType VARCHAR(50),
  productCode VARCHAR(50),
  clientCode VARCHAR(50),
  dailyVolume VARCHAR(50),



)