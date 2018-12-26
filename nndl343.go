package main
/* ゼロから作るDeepLearning 3.4.3　実装のまとめ より */

import (
  "github.com/nakamkaz/np00"
  "fmt"
  )

func InitNetwork() np00.NNetwork {

          nn := make(map[string]np00.NParray)

          nn["W1"] = np00.NParray{
                  []float64{0.1, 0.3, 0.5},
                  []float64{0.2, 0.4, 0.6},
          }
          nn["B1"] = np00.NParray{
                  []float64{0.1, 0.2, 0.3},
          }
          nn["W2"] = np00.NParray{
                  []float64{0.1, 0.4},
                  []float64{0.2, 0.5},
                  []float64{0.3, 0.6},
          }
          nn["B2"] = np00.NParray{
                  []float64{0.1, 0.2},
          }
          nn["W3"] = np00.NParray{
                  []float64{0.1, 0.3},
                  []float64{0.2, 0.4},
          }
          nn["B3"] = np00.NParray{
                  []float64{0.1, 0.2},
          }
          return nn
  }


func Forward(nn np00.NNetwork, npa np00.NParray) (npr np00.NParray) {
          W1, W2, W3 := nn["W1"], nn["W2"], nn["W3"]
          b1, b2, b3 := nn["B1"], nn["B2"], nn["B3"]
          a1 := np00.Add(np00.Dot(npa, W1), b1)
          z1 := np00.Sigmoid(a1)
          fmt.Println("z1: ",z1)
          a2 := np00.Add(np00.Dot(z1, W2), b2)
          z2 := np00.Sigmoid(a2)
          fmt.Println("z2: ",z2)
          fmt.Println("W3: ",W3)
          _ = b3
          a3 := np00.Add(np00.Dot(z2, W3), b3)
          y := np00.IdentityFunction(a3)

          return y
  }

func main(){
nn := InitNetwork()
X := np00.NParray{
        []float64{1.0,0.5},
}

Y := Forward(nn,X)
fmt.Println(Y)
}

