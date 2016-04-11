package global

import (
	"math"

	"github.com/ready-steady/adapt/algorithm/external"
	"github.com/ready-steady/adapt/basis/polynomial"
	"github.com/ready-steady/adapt/grid/equidistant"
)

type fixture struct {
	target external.Target

	surrogate *external.Surrogate

	points []float64
	values []float64
}

func prepare(fixture *fixture) *Interpolator {
	const (
		minLevel      = 1
		maxLevel      = 10
		absoluteError = 1e-6
		relativeError = 1e-3
	)

	ni, no := fixture.surrogate.Inputs, fixture.surrogate.Outputs

	grid := equidistant.NewClosed(ni)
	basis := polynomial.NewClosed(ni, 1)
	strategy := NewStrategy(ni, no, minLevel, maxLevel, absoluteError, relativeError, grid)
	interpolator := New(ni, no, grid, basis, strategy)

	return interpolator
}

var fixtureBranin = fixture{
	target: func(point, value []float64) {
		x, y := 15.0*point[0]-5.0, 15.0*point[1]
		z := 5.0/math.Pi*x - 5.1/(4.0*math.Pi*math.Pi)*x*x + y - 6.0
		value[0] = z*z + 10.0*(1.0-1.0/(8.0*math.Pi))*math.Cos(x) + 10.0
	},

	surrogate: &external.Surrogate{
		Inputs:  2,
		Outputs: 1,
		Nodes:   377,
	},

	points: []float64{
		0.0, 0.0,
		0.0, 0.1,
		0.0, 0.2,
		0.0, 0.3,
		0.0, 0.4,
		0.0, 0.5,
		0.0, 0.6,
		0.0, 0.7,
		0.0, 0.8,
		0.0, 0.9,
		0.0, 1.0,
		0.1, 0.0,
		0.1, 0.1,
		0.1, 0.2,
		0.1, 0.3,
		0.1, 0.4,
		0.1, 0.5,
		0.1, 0.6,
		0.1, 0.7,
		0.1, 0.8,
		0.1, 0.9,
		0.1, 1.0,
		0.2, 0.0,
		0.2, 0.1,
		0.2, 0.2,
		0.2, 0.3,
		0.2, 0.4,
		0.2, 0.5,
		0.2, 0.6,
		0.2, 0.7,
		0.2, 0.8,
		0.2, 0.9,
		0.2, 1.0,
		0.3, 0.0,
		0.3, 0.1,
		0.3, 0.2,
		0.3, 0.3,
		0.3, 0.4,
		0.3, 0.5,
		0.3, 0.6,
		0.3, 0.7,
		0.3, 0.8,
		0.3, 0.9,
		0.3, 1.0,
		0.4, 0.0,
		0.4, 0.1,
		0.4, 0.2,
		0.4, 0.3,
		0.4, 0.4,
		0.4, 0.5,
		0.4, 0.6,
		0.4, 0.7,
		0.4, 0.8,
		0.4, 0.9,
		0.4, 1.0,
		0.5, 0.0,
		0.5, 0.1,
		0.5, 0.2,
		0.5, 0.3,
		0.5, 0.4,
		0.5, 0.5,
		0.5, 0.6,
		0.5, 0.7,
		0.5, 0.8,
		0.5, 0.9,
		0.5, 1.0,
		0.6, 0.0,
		0.6, 0.1,
		0.6, 0.2,
		0.6, 0.3,
		0.6, 0.4,
		0.6, 0.5,
		0.6, 0.6,
		0.6, 0.7,
		0.6, 0.8,
		0.6, 0.9,
		0.6, 1.0,
		0.7, 0.0,
		0.7, 0.1,
		0.7, 0.2,
		0.7, 0.3,
		0.7, 0.4,
		0.7, 0.5,
		0.7, 0.6,
		0.7, 0.7,
		0.7, 0.8,
		0.7, 0.9,
		0.7, 1.0,
		0.8, 0.0,
		0.8, 0.1,
		0.8, 0.2,
		0.8, 0.3,
		0.8, 0.4,
		0.8, 0.5,
		0.8, 0.6,
		0.8, 0.7,
		0.8, 0.8,
		0.8, 0.9,
		0.8, 1.0,
		0.9, 0.0,
		0.9, 0.1,
		0.9, 0.2,
		0.9, 0.3,
		0.9, 0.4,
		0.9, 0.5,
		0.9, 0.6,
		0.9, 0.7,
		0.9, 0.8,
		0.9, 0.9,
		0.9, 1.0,
		1.0, 0.0,
		1.0, 0.1,
		1.0, 0.2,
		1.0, 0.3,
		1.0, 0.4,
		1.0, 0.5,
		1.0, 0.6,
		1.0, 0.7,
		1.0, 0.8,
		1.0, 0.9,
		1.0, 1.0,
	},

	values: []float64{
		3.0812909601160663e+02,
		2.5881701636202388e+02,
		2.1400493671244101e+02,
		1.7369285706285817e+02,
		1.3788077741327527e+02,
		1.0656869776369244e+02,
		7.9756618114109585e+01,
		5.7444538464526737e+01,
		3.9632458814943881e+01,
		2.6320379165361025e+01,
		1.7508299515778170e+01,
		1.7400769035750577e+02,
		1.3679889062181547e+02,
		1.0409009088612515e+02,
		7.5881291150434848e+01,
		5.2172491414744542e+01,
		3.2963691679054236e+01,
		1.8254891943363937e+01,
		8.0460922076736310e+00,
		2.3372924719833277e+00,
		1.1284927362930244e+00,
		4.4196930006027211e+00,
		1.0009094705558034e+02,
		7.3241436360338838e+01,
		5.0891925665097354e+01,
		3.3042414969855869e+01,
		1.9692904274614378e+01,
		1.0843393579372886e+01,
		6.4938828841313967e+00,
		6.6443721888899070e+00,
		1.1294861493648417e+01,
		2.0445350798406928e+01,
		3.4095840103165436e+01,
		6.5049198045714334e+01,
		4.6814985517477922e+01,
		3.3080772989241503e+01,
		2.3846560461005087e+01,
		1.9112347932768671e+01,
		1.8878135404532252e+01,
		2.3143922876295839e+01,
		3.1909710348059420e+01,
		4.5175497819823008e+01,
		6.2941285291586588e+01,
		8.5207072763350169e+01,
		3.5778175704962486e+01,
		2.4415270470287407e+01,
		1.7552365235612324e+01,
		1.5189460000937242e+01,
		1.7326554766262163e+01,
		2.3963649531587087e+01,
		3.5100744296912005e+01,
		5.0737839062236930e+01,
		7.0874933827561847e+01,
		9.5512028592886764e+01,
		1.2464912335821168e+02,
		1.0307908486409694e+01,
		4.0723196718522079e+00,
		2.3367308572947225e+00,
		5.1011420427372389e+00,
		1.2365553228179754e+01,
		2.4129964413622268e+01,
		4.0394375599064787e+01,
		6.1158786784507299e+01,
		8.6423197969949811e+01,
		1.1618760915539234e+02,
		1.5045202034083485e+02,
		6.6162059306544840e+00,
		3.7639426627708543e+00,
		5.4116793948872246e+00,
		1.1559416127003590e+01,
		2.2207152859119958e+01,
		3.7354889591236329e+01,
		5.7002626323352693e+01,
		8.1150363055469057e+01,
		1.0979809978758543e+02,
		1.4294583651970177e+02,
		1.8059357325181816e+02,
		1.8137157493546802e+01,
		1.6924228898893290e+01,
		2.0211300304239778e+01,
		2.7998371709586266e+01,
		4.0285443114932754e+01,
		5.7072514520279242e+01,
		7.8359585925625737e+01,
		1.0414665733097222e+02,
		1.3443372873631870e+02,
		1.6922080014166519e+02,
		2.0850787154701169e+02,
		1.8653238952862246e+01,
		1.7335654157995116e+01,
		2.0518069363127985e+01,
		2.8200484568260848e+01,
		4.0382899773393710e+01,
		5.7065314978526580e+01,
		7.8247730183659442e+01,
		1.0393014538879231e+02,
		1.3411256059392520e+02,
		1.6879497579905802e+02,
		2.0797739100419091e+02,
		7.4789214155018051e+00,
		4.3126895469773121e+00,
		5.6464576784528191e+00,
		1.1480225809928326e+01,
		2.1813993941403833e+01,
		3.6647762072879345e+01,
		5.5981530204354847e+01,
		7.9815298335830349e+01,
		1.0814906646730581e+02,
		1.4098283459878132e+02,
		1.7831660273025685e+02,
		1.0960889035651505e+01,
		4.2020192200259103e+00,
		1.9431494044003159e+00,
		4.1842795887747215e+00,
		1.0925409773149127e+01,
		2.2166539957523533e+01,
		3.7907670141897938e+01,
		5.8148800326272344e+01,
		8.2889930510646749e+01,
		1.1213106069502115e+02,
		1.4587219087939556e+02,
	},
}
