package main

// this function will run whenever we recieved a request
// in '/transactions' api. it will update current statitics.
func UpdateDatabase(amount float32) {
	// updating the maximum value
	if database.max < amount {
		database.max = amount
	}

	if database.min > amount {
		database.min = amount
	}

	// updating the count
	database.count++

	database.sum += amount

	database.average = database.sum / database.count

}

// as we only need avg, sum, min, max and count
// in the last 60 seconds, we need to reduce the amount
// which is added 60seconds before from current statitics
// this function refresh current statics.
func refreshDatabase(amount float32) {
	database.count--
	database.sum = database.sum - amount
	database.average = database.sum / database.count
	var min, max float32 = 999999, -99999999
	temp := queue.front
	for temp != nil {
		if temp.data < min {
			min = temp.data
		}
		if temp.data > max {
			max = temp.data
		}
		temp = temp.next
	}
	database.min = min
	database.max = max
}

func resetDB() {
	database.count = 0
	database.max = -9999999
	database.min = 999999
	database.average = 0
	database.sum = 0
}

func resetQueue() {
	queue = Queue{nil, nil}
}
