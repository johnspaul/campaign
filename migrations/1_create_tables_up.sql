
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
  dailyVolume VARCHAR(50)



);

  CREATE TABLE "campaign"
  (
    campaignId VARCHAR(50) NOT NULL PRIMARY KEY ,
    productId VARCHAR(50),
    status INTEGER,
    approvedOn INTEGER,
    approvedBy INTEGER,
    startDate INTEGER,
    endDate INTEGER,
    activeStartHour INTEGER,
    activeStartMinute INTEGER,
    activeHours INTEGER,
    executionFrequency INTEGER,
    pollingInterval INTEGER,
    lastExecutionTime INTEGER,
    currentExecutionCycleStartTime INTEGER,
    currentTargetVolume INTEGER
  );

CREATE TABLE campaign(

  campaignId varchar(36) ,
  productId varchar(36),
  status BIGINT DEFAULT 0,
  startdate  BIGINT,
  enddate BIGINT,

  activeStartHour integer,
  activeStartMinute integer,
  activeHours integer,

  pollingInterval integer,
  executionFrequency integer,
  currentExecutionCycleStartTime BIGINT,
  lastExecutionTime BIGINT,
  currentTargetVolume integer,

  approvedOn BIGINT,
  approvedBy varchar(36) ,
  PRIMARY KEY (campaignId)
);

