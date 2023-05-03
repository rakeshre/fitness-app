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
                         "instructorname" varchar NOT NULL,
                         "regstatus" varchar DEFAULT 'Available',
                         "name" varchar DEFAULT 'Class description',
                         "classtype" varchar DEFAULT 'weekly',
                         "cost" int NOT NULL,
                         "scheduleid" bigserial NOT NULL
);

CREATE TABLE "schedule" (
                            "id" BIGSERIAL PRIMARY KEY,
                            "locationid" bigserial NOT NULL,
                            "startdate" date NOT NULL,
                            "enddate" date NOT NULL,
                            "starttime" time NOT NULL,
                            "endtime" time NOT NULL,
                            "day" varchar NOT NULL
);

CREATE TABLE "membershiptypes" (
                                   "id" BIGSERIAL PRIMARY KEY,
                                   "member_type" int NOT NULL DEFAULT 0,
                                   "cost" int NOT NULL DEFAULT 0
);

CREATE TABLE "membership" (
                              "membershipid" bigserial NOT NULL,
                              "userid" bigserial NOT NULL,
                              "startdate" timestamptz NOT NULL DEFAULT (now()),
                              "expirydate" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "classcatalogue" (
                                  "id" BIGSERIAL PRIMARY KEY,
                                  "userid" bigserial NOT NULL,
                                  "courseid" bigserial NOT NULL,
                                  "enrolmentdate" timestamptz DEFAULT (now())
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

CREATE TABLE "checkinrecords" (
                                  "id" BIGSERIAL PRIMARY KEY,
                                  "userid" bigserial NOT NULL,
                                  "type" int NOT NULL DEFAULT 1,
                                  "time" timestamptz NOT NULL DEFAULT (now()),
                                  "employeeid" bigserial NOT NULL,
                                  "locationid" bigserial NOT NULL
);

CREATE TABLE "activityrecords" (
                                   "id" BIGSERIAL PRIMARY KEY,
                                   "userid" bigserial NOT NULL,
                                   "type" int NOT NULL DEFAULT 1,
                                   "time" timestamptz NOT NULL DEFAULT (now()),
                                   "locationid" bigserial NOT NULL,
                                   "deviceid" bigserial NOT NULL
);

CREATE TABLE "useractivity" (
                                "id" BIGSERIAL PRIMARY KEY,
                                "start" timestamptz NOT NULL,
                                "end" timestamptz NOT NULL,
                                "userid" bigserial NOT NULL,
                                "deviceid" bigserial NOT NULL,
                                "locationid" bigserial NOT NULL
);

CREATE TABLE "device" (
                          "id" BIGSERIAL PRIMARY KEY,
                          "description" varchar NOT NULL,
                          "status" varchar NOT NULL DEFAULT 'Free'
);

CREATE UNIQUE INDEX ON "users" ("name");

CREATE UNIQUE INDEX ON "employee" ("name");

CREATE UNIQUE INDEX ON "classcatalogue" ("userid", "courseid");

COMMENT ON COLUMN "class"."classtype" IS 'weekly daily or monthly';

COMMENT ON COLUMN "device"."status" IS 'Free,busy,not working';

ALTER TABLE "class" ADD FOREIGN KEY ("scheduleid") REFERENCES "schedule" ("id");

ALTER TABLE "membership" ADD FOREIGN KEY ("membershipid") REFERENCES "membershiptypes" ("id");

ALTER TABLE "membership" ADD FOREIGN KEY ("userid") REFERENCES "users" ("id");

ALTER TABLE "classcatalogue" ADD FOREIGN KEY ("userid") REFERENCES "users" ("id");

ALTER TABLE "classcatalogue" ADD FOREIGN KEY ("courseid") REFERENCES "class" ("id");

ALTER TABLE "checkinactivity" ADD FOREIGN KEY ("userid") REFERENCES "users" ("id");

ALTER TABLE "useractivity" ADD FOREIGN KEY ("userid") REFERENCES "users" ("id");

ALTER TABLE "useractivity" ADD FOREIGN KEY ("deviceid") REFERENCES "device" ("id");

ALTER TABLE "checkinactivity" ADD FOREIGN KEY ("locationid") REFERENCES "location" ("id");

ALTER TABLE "schedule" ADD FOREIGN KEY ("locationid") REFERENCES "location" ("id");

ALTER TABLE "checkinactivity" ADD FOREIGN KEY ("employeeid") REFERENCES "employee" ("id");

ALTER TABLE "employee" ADD FOREIGN KEY ("locationid") REFERENCES "location" ("id");
