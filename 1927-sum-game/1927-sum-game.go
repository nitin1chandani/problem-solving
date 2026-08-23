func sumGame(num string) bool {
    var noOfQuestionMarksOnLeftSide, noOfQuestionMarksOnRightSide, knownSumOfLeftSide, knownSumOfRightSide int
    n := len(num)

    for i, v := range num{
        if v == '?'{
            if i < n/2{
                noOfQuestionMarksOnLeftSide++
            }else{
                noOfQuestionMarksOnRightSide++
            }
        }else{
            if i<n/2{
                knownSumOfLeftSide += int(v - '0')
            }else{
                knownSumOfRightSide += int(v - '0')
            }
        }
    }

    totalQnMarks := noOfQuestionMarksOnLeftSide + noOfQuestionMarksOnRightSide

    if totalQnMarks % 2 == 1 {
        return true
    }

    leftSum := 2*knownSumOfLeftSide + 9*noOfQuestionMarksOnLeftSide
    rightSum := 2*knownSumOfRightSide + 9*noOfQuestionMarksOnRightSide

    if leftSum==rightSum{
        return false 
    } 
    return true
}