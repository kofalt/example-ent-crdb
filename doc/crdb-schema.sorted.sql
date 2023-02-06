
-- show create table users
-- show create table vehicles
-- show create table rides
-- show create table vehicle_location_histories
-- show create table user_promo_codes

CREATE TABLE public.users (
	address VARCHAR NULL,
	city VARCHAR NOT NULL,
	credit_card VARCHAR NULL,
	id UUID NOT NULL,
	name VARCHAR NULL,
	CONSTRAINT users_pkey PRIMARY KEY (city ASC, id ASC)
)

CREATE TABLE public.vehicles (
	city VARCHAR NOT NULL,
	creation_time TIMESTAMP NULL,
	current_location VARCHAR NULL,
	ext JSONB NULL,
	id UUID NOT NULL,
	owner_id UUID NULL,
	status VARCHAR NULL,
	type VARCHAR NULL,
	CONSTRAINT vehicles_pkey PRIMARY KEY (city ASC, id ASC),
	CONSTRAINT vehicles_city_owner_id_fkey FOREIGN KEY (city, owner_id) REFERENCES public.users(city, id),
	INDEX vehicles_auto_index_fk_city_ref_users (city ASC, owner_id ASC)
)

CREATE TABLE public.rides (
	city VARCHAR NOT NULL,
	end_address VARCHAR NULL,
	end_time TIMESTAMP NULL,
	id UUID NOT NULL,
	revenue DECIMAL(10,2) NULL,
	rider_id UUID NULL,
	start_address VARCHAR NULL,
	start_time TIMESTAMP NULL,
	vehicle_city VARCHAR NULL,
	vehicle_id UUID NULL,
	CONSTRAINT rides_pkey PRIMARY KEY (city ASC, id ASC),
	CONSTRAINT rides_city_rider_id_fkey FOREIGN KEY (city, rider_id) REFERENCES public.users(city, id),
	CONSTRAINT rides_vehicle_city_vehicle_id_fkey FOREIGN KEY (vehicle_city, vehicle_id) REFERENCES public.vehicles(city, id),
	INDEX rides_auto_index_fk_city_ref_users (city ASC, rider_id ASC),
	INDEX rides_auto_index_fk_vehicle_city_ref_vehicles (vehicle_city ASC, vehicle_id ASC),
	CONSTRAINT check_vehicle_city_city CHECK (vehicle_city = city)
)

CREATE TABLE public.vehicle_location_histories (
	"timestamp" TIMESTAMP NOT NULL,
	city VARCHAR NOT NULL,
	lat FLOAT8 NULL,
	long FLOAT8 NULL,
	ride_id UUID NOT NULL,
	CONSTRAINT vehicle_location_histories_pkey PRIMARY KEY (city ASC, ride_id ASC, "timestamp" ASC),
	CONSTRAINT vehicle_location_histories_city_ride_id_fkey FOREIGN KEY (city, ride_id) REFERENCES public.rides(city, id)
)

CREATE TABLE public.user_promo_codes (
	"timestamp" TIMESTAMP NULL,
	city VARCHAR NOT NULL,
	code VARCHAR NOT NULL,
	usage_count INT8 NULL,
	user_id UUID NOT NULL,
	CONSTRAINT user_promo_codes_pkey PRIMARY KEY (city ASC, user_id ASC, code ASC),
	CONSTRAINT user_promo_codes_city_user_id_fkey FOREIGN KEY (city, user_id) REFERENCES public.users(city, id)
)

