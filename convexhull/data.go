package main

// Point represent basic point struct (on 2 Dimensional lane)
type Point struct {
	X float64
	Y float64
}

// GenerateTestData returns a test data for initial points of convex hull
func GenerateTestData() []Point {
	return []Point{
		{X: 0.3215348546593775, Y: 0.03629583077160248},
		{X: 0.02402358131857918, Y: -0.2356728797179394},
		{X: 0.04590851212470659, Y: -0.4156409924995536},
		{X: 0.3218384001607433, Y: 0.1379850698988746},
		{X: 0.11506479756447, Y: -0.1059521474930943},
		{X: 0.2622539999543261, Y: -0.29702873322836},
		{X: -0.161920957418085, Y: -0.4055339716426413},
		{X: 0.1905378631228002, Y: 0.3698601009043493},
		{X: 0.2387090918968516, Y: -0.01629827079949742},
		{X: 0.07495888748668034, Y: -0.1659825110491202},
		{X: 0.3319341836794598, Y: -0.1821814101954749},
		{X: 0.07703635755650362, Y: -0.2499430638271785},
		{X: 0.2069242999022122, Y: -0.2232970760420869},
		{X: 0.04604079532068295, Y: -0.1923573186549892},
		{X: 0.05054295812784038, Y: 0.4754929463150845},
		{X: -0.3900589168910486, Y: 0.2797829520700341},
		{X: 0.3120693385713448, Y: -0.0506329867529059},
		{X: 0.01138812723698857, Y: 0.4002504701728471},
		{X: 0.009645149586391732, Y: 0.1060251100976254},
		{X: -0.03597933197019559, Y: 0.2953639456959105},
		{X: 0.1818290866742182, Y: 0.001454397571696298},
		{X: 0.444056063372694, Y: 0.2502497166863175},
		{X: -0.05301752458607545, Y: -0.06553921621808712},
		{X: 0.4823896228171788, Y: -0.4776170002088109},
		{X: -0.3089226845734964, Y: -0.06356112199235814},
		{X: -0.271780741188471, Y: 0.1810810595574612},
		{X: 0.4293626522918815, Y: 0.2980897964891882},
		{X: -0.004796652127799228, Y: 0.382663812844701},
		{X: 0.430695573269106, Y: -0.2995073500084759},
		{X: 0.1799668387323309, Y: -0.2973467472915973},
		{X: 0.4932166845474547, Y: 0.4928094162538735},
		{X: -0.3521487911717489, Y: 0.4352656197131292},
		{X: -0.4907368011686362, Y: 0.1865826865533206},
		{X: -0.1047924716070224, Y: -0.247073392148198},
		{X: 0.4374961861758457, Y: -0.001606279519951237},
		{X: 0.003256207800708899, Y: -0.2729194320486108},
		{X: 0.04310378203457577, Y: 0.4452604050238248},
		{X: 0.4916198379282093, Y: -0.345391701297268},
		{X: 0.001675087028811806, Y: 0.1531837672490476},
		{X: -0.4404289572876217, Y: -0.2894855991839297},
	}
}

// GenerateCheckData returns the real convex hull points to check if the result is correct
func GenerateCheckData() []Point {
	return []Point{
		{X: -0.161920957418085, Y: -0.4055339716426413},
		{X: 0.05054295812784038, Y: 0.4754929463150845},
		{X: 0.4823896228171788, Y: -0.4776170002088109},
		{X: 0.4932166845474547, Y: 0.4928094162538735},
		{X: -0.3521487911717489, Y: 0.4352656197131292},
		{X: -0.4907368011686362, Y: 0.1865826865533206},
		{X: 0.4916198379282093, Y: -0.345391701297268},
		{X: -0.4404289572876217, Y: -0.2894855991839297},
	}
}
