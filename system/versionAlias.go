package system

import "github.com/Kiuryy/uaparser/version"

var macVersionAlias = map[version.Version]string{
	{10, 1, 0}:  "Puma",
	{10, 2, 0}:  "Jaguar",
	{10, 3, 0}:  "Panther",
	{10, 4, 0}:  "Tiger",
	{10, 5, 0}:  "Leopard",
	{10, 6, 0}:  "Snow Leopard",
	{10, 7, 0}:  "Lion",
	{10, 8, 0}:  "Mountain Lion",
	{10, 9, 0}:  "Mavericks",
	{10, 10, 0}: "Yosemite",
	{10, 11, 0}: "El Capitan",
	{10, 12, 0}: "Sierra",
	{10, 13, 0}: "High Sierra",
	{10, 14, 0}: "Mojave",
	{10, 15, 0}: "Catalina",
	{10, 16, 0}: "Big Sur",
}

var windowsVersionAlias = map[version.Version]string{
	{6, 3, 0}: "8.1",
	{6, 2, 0}: "8",
	{6, 1, 0}: "7",
	{6, 0, 0}: "Vista",
	{5, 2, 0}: "XP",
	{5, 1, 0}: "XP",
	{5, 0, 0}: "2000",
}

var androidVersionAlias = map[version.Version]string{
	{1, 5, 0}: "Cupcake",
	{1, 6, 0}: "Donut",
	{2, 0, 0}: "Éclair",
	{2, 1, 0}: "Éclair",
	{2, 2, 0}: "Froyo",
	{2, 3, 0}: "Gingerbread",
	{3, 0, 0}: "Honeycomb",
	{3, 1, 0}: "Honeycomb",
	{3, 2, 0}: "Honeycomb",
	{4, 0, 0}: "Ice Cream Sandwich",
	{4, 1, 0}: "Jelly Bean",
	{4, 2, 0}: "Jelly Bean",
	{4, 3, 0}: "Jelly Bean",
	{4, 4, 0}: "KitKat",
	{5, 0, 0}: "Lollipop",
	{5, 1, 0}: "Lollipop",
	{6, 0, 0}: "Marshmallow",
	{7, 0, 0}: "Nougat",
	{7, 1, 0}: "Nougat",
	{8, 0, 0}: "Oreo",
	{8, 1, 0}: "Oreo",
	{9, 0, 0}: "Pie",
}
