package fishing

import (
	"fmt"
	"github.com/fogleman/gg"
	"image/color"
	"math"
	"math/rand"
)

const X = 5000
const Y = 5000

var SeaMap = [X][Y]int{}

// harita oluşturulur. balık yakalanır.
func GenerateMap() {
	dc := gg.NewContext(X, Y)
	dc.DrawRectangle(0, 0, X, Y)
	dc.SetColor(color.White)
	dc.Fill()
	dc.SetColor(color.White)
	dc.Stroke()
	for i := 0; i < X; i++ {
		if math.Mod(float64(i*100)/X, 1) == 0 {
			fmt.Printf("%.2f\n", float64(i*100)/X)
		}
		for j := 0; j < Y; j++ {
			SeaMap[i][j] = 0
		}
	}
	for i := 0; i < X; i++ {
		if math.Mod(float64(i*100)/X, 1) == 0 {
			fmt.Printf("%.2f\n", float64(i*100)/X)
		}
		for j := 0; j < Y; j++ {
			rate := rand.Int() % 1001
			rate2 := rand.Int() % 100
			// 50 // 30 // 20
			if rate == 1000 && rate2 == 99 {
				if i > 4 && j > 4 && i < X-4 && j < Y-4 {
					for k := -2; k <= 2; k++ {
						for l := -2; l <= 2; l++ {
							val := int(math.Abs(math.Max(float64(k), float64(l))))
							dc.DrawPoint(float64(i-k), float64(j-l), 0.7)
							dc.SetRGB(0, 0, 100)
							dc.Stroke()
							SeaMap[i-k][j-l] = (4 - val) * 10
						}
					}
				}
			} else if rate >= 990 && rate2 >= 95 {
				if SeaMap[i][j] == 0 {
					dc.DrawLine(float64(i-1), float64(j-1), float64(i+1), float64(j+1))
					dc.SetRGB(100, 0, 0)
					dc.Stroke()
					dc.DrawLine(float64(i+1), float64(j-1), float64(i-1), float64(j+1))
					dc.SetRGB(100, 0, 0)
					dc.Stroke()
					SeaMap[i][j] = 2
				}
			} else if rate >= 970 && rate2 >= 95 {
				if SeaMap[i][j] == 0 {
					dc.DrawLine(float64(i+1), float64(j-1), float64(i-1), float64(j+1))
					dc.SetRGB(0, 100, 0)
					dc.Stroke()
					dc.DrawLine(float64(i-1), float64(j-1), float64(i+1), float64(j+1))
					dc.SetRGB(0, 100, 0)
					dc.Stroke()
					SeaMap[i][j] = 1
				}

			}
		}
	}
	dc.SavePNG("seaMap.png")
}

func FetchFish(x int, y int) int {
	return SeaMap[x][y]
}
