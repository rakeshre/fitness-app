CREATE TABLE "users" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "hashedpassword" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "employee" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "hashedpassword" varchar NOT NULL,
  "locationid" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "class" (
  "id" BIGSERIAL PRIMARY KEY,
  "instructorid" bigserial NOT NULL,
  "reg_status" varchar DEFAULT 'Open',
  "start_time" timestamptz NOT NULL,
  "end_time" timestamptz NOT NULL,
  "description" varchar DEFAULT 'Class description',
  "classtype" varchar DEFAULT 'weekly',
  "locationid" bigserial NOT NULL
);

CREATE TABLE "membership" (
  "id" BIGSERIAL,
  "userid" bigserial,
  "member_type" int NOT NULL DEFAULT 0,
  "expiry_date" timestamptz,
  PRIMARY KEY ("id", "userid")
);

CREATE TABLE "classcatalogue" (
  "id" BIGSERIAL PRIMARY KEY,
  "userid" bigserial NOT NULL,
  "courseid" bigserial NOT NULL
);

CREATE TABLE "location" (
  "id" BIGSERIAL PRIMARY KEY,
  "city" varchar NOT NULL,
  "state" varchar NOT NULL,
  "zipcode" varchar NOT NULL
);

CREATE TABLE "checkinactivity" (
  "id" BIGSERIAL PRIMARY KEY,
  "checkin" timestamptz NOT NULL,
  "checkout" timestamptz NOT NULL,
  "userid" bigserial NOT NULL,
  "employeeid" bigserial NOT NULL,
  "locationid" bigserial NOT NULL
);

CREATE TABLE "useractivity" (
  "id" BIGSERIAL PRIMARY KEY,
  "start" timestamptz NOT NULL,
  "end" timestamptz NOT NULL,
  "userid" bigserial NOT NULL,
  "deviceid" bigserial NOT NULL
);

CREATE TABLE "device" (
  "id" BIGSERIAL PRIMARY KEY,
  "description" varchar NOT NULL,
  "status" varchar NOT NULL DEFAULT 'Free'
);

CREATE UNIQUE INDEX ON "users" ("name");

CREATE UNIQUE INDEX ON "employee" ("name");

CREATE UNIQUE INDEX ON "membership" ("userid");

CREATE UNIQUE INDEX ON "classcatalogue" ("userid", "courseid");

COMMENT ON COLUMN "class"."classtype" IS 'weekly daily or monthly';

COMMENT ON COLUMN "membership"."member_type" IS '0 is admin 1 is member 2 is non member ';

COMMENT ON COLUMN "device"."status" IS 'Free,busy,not working';

ALTER TABLE "class" ADD FOREIGN KEY ("instructorid") REFERENCES "employee" ("id");

ALTER TABLE "membership" ADD FOREIGN KEY ("userid") REFERENCES "users" ("id");

ALTER TABLE "classcatalogue" ADD FOREIGN KEY ("userid") REFERENCES "users" ("id");

ALTER TABLE "classcatalogue" ADD FOREIGN KEY ("courseid") REFERENCES "class" ("id");

ALTER TABLE "checkinactivity" ADD FOREIGN KEY ("userid") REFERENCES "users" ("id");

ALTER TABLE "useractivity" ADD FOREIGN KEY ("userid") REFERENCES "users" ("id");

ALTER TABLE "useractivity" ADD FOREIGN KEY ("deviceid") REFERENCES "device" ("id");

ALTER TABLE "checkinactivity" ADD FOREIGN KEY ("locationid") REFERENCES "location" ("id");

ALTER TABLE "class" ADD FOREIGN KEY ("locationid") REFERENCES "location" ("id");

ALTER TABLE "checkinactivity" ADD FOREIGN KEY ("employeeid") REFERENCES "employee" ("id");

ALTER TABLE "employee" ADD FOREIGN KEY ("locationid") REFERENCES "location" ("id");
