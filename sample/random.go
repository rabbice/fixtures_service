package sample

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
	pb "github.com/rabbice/fixturesbook/pb"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomInt(min, max int) int {
	return min + rand.Int()%(max-min+1)
}

func randomID() string {
	return uuid.New().String()
}

func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}

func randomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randomTempUnit() *pb.Temp_Unit {
	switch rand.Intn(2) {
	case 1:
		return pb.Temp_FAHRENHEIT.Enum()
	default:
		return pb.Temp_FAHRENHEIT.Enum()
	}
}

func randomFixture(team string) string {
	if team == "home" {
		return randomStringFromSet(
			"Manchester United",
			"Tottenham",
			"Chelsea",
			"Leeds United",
			"Brighton Hove & Albion",
		)
	}

	return randomStringFromSet(
		"Arsenal",
		"Bournemouth F.C.",
		"Wolverhampton",
		"Stoke City",
		"Manchester City",
	)
}

func randomTemp() *pb.Temp {
	temp := randomFloat32(-15, 30)

	temperature := &pb.Temp{
		Temp: float32(temp),
		Unit: *randomTempUnit(),
	}
	return temperature
}

func randomTime() string {
	return randomStringFromSet("18:45:00", "19:30:00", "16:00:00", "17:45:00", "21:30:00")
}

func randomWindSpeed() string {
	return randomStringFromSet("23.5 m/s", "17.22 m/s")
}

func randomScore() string {
	return randomStringFromSet("3-2", "1-1", "0-0", "2-0", "3-1")
}

func randomDate() string {
	return randomStringFromSet(
		"2023-01-24",
		"2019-07-12",
		"2018-08-17",
		"2020-05-13",
		"2021-01-22",
	)
}

func randomOfficial() string {
	return randomStringFromSet(
		"Mehmet Yildiraz",
		"Mike Dean",
		"Howard Webb",
		"Paul Tierney",
		"Simon Hooper",
		"Anthony Taylor",
	)
}
