package sample

import (
	pb "github.com/rabbice/fixturesbook/pb"
)

func NewScore() *pb.Score {
	score := &pb.Score{
		Localteam: randomScore(),
		Awayteam:  randomScore(),
		Htscore:   randomScore(),
		Ftscore:   randomScore(),
	}
	return score
}

func NewTime() *pb.Time {
	time := &pb.Time{
		Date:     randomDate(),
		Time:     randomTime(),
		Timezone: "UTC",
	}
	return time
}

func NewFixture() *pb.Fixture {

	fixture := &pb.Fixture{
		ID:       randomID(),
		Hometeam: randomFixture("home"),
		Awayteam: randomFixture("away"),
		Score:    NewScore(),
		Time:     NewTime(),
		Official: randomOfficial(),
		Stats:    NewStats(),
		Pitch:    NewPitch(),
	}
	return fixture
}

func NewPitch() *pb.Pitch {
	pitch := &pb.Pitch{
		Name:       "Wembley Stadium",
		Commentary: true,
		Attendance: 80000,
		Weather:    NewWeatherReport(),
	}
	return pitch
}

func NewWind() *pb.Wind {
	speed := randomWindSpeed()
	degree := randomInt(20, 330)

	wind := &pb.Wind{
		Speed:  speed,
		Degree: int32(degree),
	}
	return wind

}

func NewStats() *pb.Stats {
	stats := &pb.Stats{
		Shots: []*pb.Shots{
			{Total: 34, Ongoal: 23},
		},
		Goals: []*pb.Goals{
			{Scored: 5, Assists: 5, Conceded: 5, Owngoals: 0},
		},
	}
	return stats
}

func NewWeatherReport() *pb.Weather {

	weather := &pb.Weather{
		Code:     "Clouds",
		Type:     "Scattered Clouds",
		Temp:     randomTemp(),
		Wind:     NewWind(),
		Humidity: randomStringFromSet("5%", "27%", "44%"),
	}
	return weather
}
