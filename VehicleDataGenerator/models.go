package VehicleDataGenerator

type ConnectJson struct {
	Index     int    `json:"index"`
	Type      string `json:"type"`
	Timestamp string `json:"timestamp"`
	Vin       string `json:"vin"`
}

type GpsLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Vehicle struct {
	Vin                     string      `json:"vin"`
	FrontLidarDistance      float64     `json:"front_lidar_distance"`
	FrontUltrasonicDistance float64     `json:"front_ultrasonic_distance"`
	RearUltrasonicDistance  float64     `json:"rear_ultrasonic_distance"`
	WheelSpeed              float64     `json:"wheel_speed"`
	GpsLocation             GpsLocation `json:"gps_location"`
	GpsSpeed                float64     `json:"gps_speed"`
	GpsDirection            float64     `json:"gps_direction"`
	MagnetometerDirection   float64     `json:"magnetometer_direction"`
}

type UpdateJson struct {
	Index     int     `json:"index"`
	Type      string  `json:"type"`
	Timestamp string  `json:"timestamp"`
	Vehicle   Vehicle `json:"vehicle"`
}

type NewVehicle struct {
	Vin                string  `json:"VIN"`
	Longitude          float64 `json:"Longitude"`
	Latitude           float64 `json:"Latitude"`
	CarDirection       float64 `json:"CarDirection"`
	DistanceUltrasonic float64 `json:"DistanceUltrasonic"`
	RearUltrasonic     float64 `json:"RearUltrasonic"`
	DistanceLidar      float64 `json:"DistanceLidar"`
	SpeedFrontLeft     float64 `json:"SpeedFrontLeft"`
	SpeedFrontRight    float64 `json:"SpeedFrontRight"`
	SpeedRearLeft      float64 `json:"SpeedRearLeft"`
	SpeedRearRight     float64 `json:"SpeedRearRight"`
}

type NewUpdateJson struct {
	Index     int        `json:"index"`
	Type      string     `json:"type"`
	Timestamp string     `json:"timestamp"`
	Vehicle   NewVehicle `json:"vehicle"`
}
