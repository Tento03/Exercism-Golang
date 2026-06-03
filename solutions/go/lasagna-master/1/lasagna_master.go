package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string,minute int) int{
    totalSlice:=len(layers)

    if(minute==0){
        return totalSlice*2
    }

    return totalSlice * minute
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var totalNoodles=0
    var totalSauce=0

    for _,layer := range layers{
        if (layer=="noodles"){
            totalNoodles++
        }
        if (layer=="sauce"){
            totalSauce++
        }
    }

    resultNoodles:=totalNoodles*50
    resultSauce:=float64(totalSauce)*0.2

    return resultNoodles,resultSauce
}

func AddSecretIngredient(friendList, myList []string) {
	nFriendList:=len(friendList)
    nMyList:=len(myList)
    lastIndexFriendList:=friendList[nFriendList-1]

    myList[nMyList-1]=lastIndexFriendList
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	portion:=float64(portions)/2
    result:= []float64{}

    for _,val :=  range quantities{
        result=append(result,val * portion)
    }
    return result
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
