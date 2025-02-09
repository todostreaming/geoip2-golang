package geoip2

import (
	"math/rand"
	"net"
	"testing"

	. "gopkg.in/check.v1"
)

func TestGeoIP2(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestReader(c *C) {
	reader, err := Open("test-data/test-data/GeoIP2-City-Test.mmdb")
	if err != nil {
		c.Log(err)
		c.Fail()
	}
	defer reader.Close()

	record, err := reader.City(net.ParseIP("81.2.69.160"))
	if err != nil {
		c.Log(err)
		c.Fail()
	}

	m := reader.Metadata()
	c.Assert(m.BinaryFormatMajorVersion, Equals, uint(2))
	c.Assert(m.BinaryFormatMinorVersion, Equals, uint(0))
	c.Assert(m.BuildEpoch, Equals, uint(1436981935))
	c.Assert(m.DatabaseType, Equals, "GeoIP2-City")
	c.Assert(m.Description, DeepEquals, map[string]string{
		"en": "GeoIP2 City Test Database (a small sample of real GeoIP2 data)",
		"zh": "小型数据库",
	})
	c.Assert(m.IPVersion, Equals, uint(6))
	c.Assert(m.Languages, DeepEquals, []string{"en", "zh"})
	c.Assert(m.NodeCount, Equals, uint(1240))
	c.Assert(m.RecordSize, Equals, uint(28))

	c.Assert(record.City.GeoNameID, Equals, uint(2643743))
	c.Assert(record.City.Names, DeepEquals, map[string]string{
		"de":    "London",
		"en":    "London",
		"es":    "Londres",
		"fr":    "Londres",
		"ja":    "ロンドン",
		"pt-BR": "Londres",
		"ru":    "Лондон",
	})
	c.Assert(record.Continent.GeoNameID, Equals, uint(6255148))
	c.Assert(record.Continent.Code, Equals, "EU")
	c.Assert(record.Continent.Names, DeepEquals, map[string]string{
		"de":    "Europa",
		"en":    "Europe",
		"es":    "Europa",
		"fr":    "Europe",
		"ja":    "ヨーロッパ",
		"pt-BR": "Europa",
		"ru":    "Европа",
		"zh-CN": "欧洲",
	})

	c.Assert(record.Country.GeoNameID, Equals, uint(2635167))
	c.Assert(record.Country.IsoCode, Equals, "GB")
	c.Assert(record.Country.Names, DeepEquals, map[string]string{
		"de":    "Vereinigtes Königreich",
		"en":    "United Kingdom",
		"es":    "Reino Unido",
		"fr":    "Royaume-Uni",
		"ja":    "イギリス",
		"pt-BR": "Reino Unido",
		"ru":    "Великобритания",
		"zh-CN": "英国",
	})

	c.Assert(record.Location.Latitude, Equals, 51.5142)
	c.Assert(record.Location.Longitude, Equals, -0.0931)
	c.Assert(record.Location.TimeZone, Equals, "Europe/London")

	c.Assert(record.Subdivisions[0].GeoNameID, Equals, uint(6269131))
	c.Assert(record.Subdivisions[0].IsoCode, Equals, "ENG")
	c.Assert(record.Subdivisions[0].Names, DeepEquals, map[string]string{
		"en":    "England",
		"pt-BR": "Inglaterra",
		"fr":    "Angleterre",
		"es":    "Inglaterra",
	})

	c.Assert(record.RegisteredCountry.GeoNameID, Equals, uint(6252001))
	c.Assert(record.RegisteredCountry.IsoCode, Equals, "US")
	c.Assert(record.RegisteredCountry.Names, DeepEquals, map[string]string{
		"de":    "USA",
		"en":    "United States",
		"es":    "Estados Unidos",
		"fr":    "États-Unis",
		"ja":    "アメリカ合衆国",
		"pt-BR": "Estados Unidos",
		"ru":    "США",
		"zh-CN": "美国",
	})
}

func (s *MySuite) TestMetroCode(c *C) {
	reader, err := Open("test-data/test-data/GeoIP2-City-Test.mmdb")
	if err != nil {
		c.Log(err)
		c.Fail()
	}
	defer reader.Close()

	record, err := reader.City(net.ParseIP("216.160.83.56"))
	if err != nil {
		c.Log(err)
		c.Fail()
	}

	c.Assert(record.Location.MetroCode, Equals, uint(819))
}

func (s *MySuite) TestConnectionType(c *C) {
	reader, err := Open("test-data/test-data/GeoIP2-Connection-Type-Test.mmdb")
	if err != nil {
		c.Log(err)
		c.Fail()
	}
	defer reader.Close()

	record, err := reader.ConnectionType(net.ParseIP("1.0.1.0"))
	if err != nil {
		c.Log(err)
		c.Fail()
	}
	c.Assert(record.ConnectionType, Equals, "Cable/DSL")

}

func (s *MySuite) TestDomain(c *C) {
	reader, err := Open("test-data/test-data/GeoIP2-Domain-Test.mmdb")
	if err != nil {
		c.Log(err)
		c.Fail()
	}
	defer reader.Close()

	record, err := reader.Domain(net.ParseIP("1.2.0.0"))
	if err != nil {
		c.Log(err)
		c.Fail()
	}
	c.Assert(record.Domain, Equals, "maxmind.com")

}

func (s *MySuite) TestISP(c *C) {
	reader, err := Open("test-data/test-data/GeoIP2-ISP-Test.mmdb")
	if err != nil {
		c.Log(err)
		c.Fail()
	}
	defer reader.Close()

	record, err := reader.ISP(net.ParseIP("1.128.0.0"))
	if err != nil {
		c.Log(err)
		c.Fail()
	}

	c.Assert(record.AutonomousSystemNumber, Equals, uint(1221))

	c.Assert(record.AutonomousSystemOrganization, Equals, "Telstra Pty Ltd")
	c.Assert(record.ISP, Equals, "Telstra Internet")
	c.Assert(record.Organization, Equals, "Telstra Internet")

}

func (s *MySuite) TestAnonymousIP(c *C) {
	reader, err := Open("test-data/test-data/GeoIP2-Anonymous-IP-Test.mmdb")
	if err != nil {
		c.Log(err)
		c.Fail()
	}
	defer reader.Close()

	record, err := reader.AnonymousIP(net.ParseIP("1.2.0.0"))
	if err != nil {
		c.Log(err)
		c.Fail()
	}

	c.Assert(record.IsAnonymous, Equals, true)

	c.Assert(record.IsAnonymousVPN, Equals, true)
	c.Assert(record.IsHostingProvider, Equals, false)
	c.Assert(record.IsPublicProxy, Equals, false)
	c.Assert(record.IsTorExitNode, Equals, false)

}

// This ensures the compiler does not optimize away the function call
var cityResult *City

func BenchmarkMaxMindDB(b *testing.B) {
	db, err := Open("GeoLite2-City.mmdb")
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	r := rand.New(rand.NewSource(0))

	var city *City

	for i := 0; i < b.N; i++ {
		ip := randomIPv4Address(b, r)
		city, err = db.City(ip)
		if err != nil {
			b.Fatal(err)
		}
	}
	cityResult = city
}

func randomIPv4Address(b *testing.B, r *rand.Rand) net.IP {
	num := r.Uint32()
	ip := []byte{byte(num >> 24), byte(num >> 16), byte(num >> 8),
		byte(num)}

	return ip
}
