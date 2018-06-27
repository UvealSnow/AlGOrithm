package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // Seed the random number generator with current unix nano timestamp.
    rand.Seed(time.Now().UnixNano())

    // Define an array with 5 random integers.
    var unsorted_array = randomArray(50)

    // Print the array we get.
    fmt.Println("Unsorted array:", unsorted_array)

    // Sort the array.
    var sorted_array = sortArray(unsorted_array, 0)

    // Print sorted array.
    fmt.Println("Sorted array:", sorted_array)
}

 /*
  * Gets a random integer using the std library.
  *
  * @param  int min
  * @param  int max
  * @return int
  */
func randomInteger(min int, max int) int {
    return min + rand.Intn(max - min)
}

/*
 * Gets an array made of random integers of a given length.
 *
 * @param  int size
 * @return []int array
 */
func randomArray(size int) []int {
    array := make([]int, size)

    for i := 0; i < size; i++ {
        array[i] = randomInteger(1, 100)
    }

    return array
}

/*
 * Sorts the given array using tail recursion.
 * Sorts the array using insertion sort algorithm.
 *
 * @param  []int array
 * @param  int pivot
 * @return []int array
 */
func sortArray(array []int, pivot int) []int {
    // Check if pivot is still within array boundaries.
    if (pivot < len(array)) {
        // Loop over all of the arrays' keys.
        for i := pivot + 1; i < len(array); i++ {
            // If array[pivot] > array[i] positions must be swapped.
            if (array[pivot] > array[i]) {
                var tmp = array[i]

                array[i] = array[pivot]
                array[pivot] = tmp
            }
        }

        // Iterate the function with new pivot
        return sortArray(array, pivot + 1)
    }

    // Array is sorted.
    return array
}
