
-- show create table users
-- show create table vehicles
-- show create table rides
-- show create table vehicle_location_histories
-- show create table user_promo_codes

CREATE TABLE public.users (
	id UUID NOT NULL,
	city VARCHAR NOT NULL,
	name VARCHAR NULL,
	address VARCHAR NULL,
	credit_card VARCHAR NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id ASC),
	INDEX user_city_id (city ASC, id ASC)
)

CREATE TABLE public.vehicles (
	id UUID NOT NULL,
	city VARCHAR NOT NULL,
	type VARCHAR NULL,
	creation_time TIMESTAMP NULL,
	status VARCHAR NULL,
	current_location VARCHAR NULL,
	ext JSONB NULL,
	owner_id UUID NULL,
	CONSTRAINT vehicles_pkey PRIMARY KEY (id ASC),
	CONSTRAINT vehicles_users_vehicles FOREIGN KEY (owner_id) REFERENCES public.users(id) ON DELETE SET NULL,
	INDEX vehicle_city_id (city ASC, id ASC),
	INDEX vehicle_city_owner_id (city ASC, owner_id ASC)
)

CREATE TABLE public.rides (
	id UUID NOT NULL,
	city VARCHAR NOT NULL,
	vehicle_city VARCHAR NULL,
	start_address VARCHAR NULL,
	end_address VARCHAR NULL,
	start_time TIMESTAMP NULL,
	end_time TIMESTAMP NULL,
	revenue DECIMAL(10,2) NULL,
	rider_id UUID NULL,
	vehicle_id UUID NULL,
	CONSTRAINT rides_pkey PRIMARY KEY (id ASC),
	CONSTRAINT rides_users_user FOREIGN KEY (rider_id) REFERENCES public.users(id) ON DELETE SET NULL,
	CONSTRAINT rides_vehicles_rides FOREIGN KEY (vehicle_id) REFERENCES public.vehicles(id) ON DELETE SET NULL,
	INDEX ride_city_id (city ASC, id ASC)
)

CREATE TABLE public.vehicle_location_histories (
	id UUID NOT NULL,
	city VARCHAR NOT NULL,
	"timestamp" TIMESTAMP NOT NULL,
	lat FLOAT8 NULL,
	long FLOAT8 NULL,
	ride_id UUID NOT NULL,
	CONSTRAINT vehicle_location_histories_pkey PRIMARY KEY (id ASC),
	CONSTRAINT vehicle_location_histories_rides_vehicle_location_histories FOREIGN KEY (ride_id) REFERENCES public.rides(id),
	INDEX vehiclelocationhistory_city_ride_id_timestamp (city ASC, ride_id ASC, "timestamp" ASC)
)

CREATE TABLE public.user_promo_codes (
	id UUID NOT NULL,
	city VARCHAR NOT NULL,
	code VARCHAR NOT NULL,
	"timestamp" TIMESTAMP NULL,
	usage_count INT8 NULL,
	user_id UUID NOT NULL,
	CONSTRAINT user_promo_codes_pkey PRIMARY KEY (id ASC),
	CONSTRAINT user_promo_codes_users_user_promo_codes FOREIGN KEY (user_id) REFERENCES public.users(id),
	INDEX userpromocode_city_user_id_code (city ASC, user_id ASC, code ASC)
)
