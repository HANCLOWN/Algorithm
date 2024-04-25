package main

import "math"

//2739. 总行驶距离
//提示
//卡车有两个油箱。给你两个整数，mainTank 表示主油箱中的燃料（以升为单位），additionalTank 表示副油箱中的燃料（以升为单位）。
//该卡车每耗费 1 升燃料都可以行驶 10 km。每当主油箱使用了 5 升燃料时，如果副油箱至少有 1 升燃料，则会将 1 升燃料从副油箱转移到主油箱。
//返回卡车可以行驶的最大距离。
//注意：从副油箱向主油箱注入燃料不是连续行为。这一事件会在每消耗 5 升燃料时突然且立即发生。
//示例 1：
//输入：mainTank = 5, additionalTank = 10
//输出：60
//解释：
//在用掉 5 升燃料后，主油箱中燃料还剩下 (5 - 5 + 1) = 1 升，行驶距离为 50km 。
//在用掉剩下的 1 升燃料后，没有新的燃料注入到主油箱中，主油箱变为空。
//总行驶距离为 60km 。
//示例 2：
//输入：mainTank = 1, additionalTank = 2
//输出：10
//解释：
//在用掉 1 升燃料后，主油箱变为空。
//总行驶距离为 10km 。
//提示：
//1 <= mainTank, additionalTank <= 100

func distanceTraveled(mainTank int, additionalTank int) int {
	totalDistance := 0
	for mainTank >= 5 {
		// 消耗主油箱的5升燃料行驶50km
		totalDistance += 50
		mainTank -= 5

		// 如果主油箱消耗完5升并且副油箱还有燃料，从副油箱转移1升燃料到主油箱
		if additionalTank > 0 {
			mainTank += 1
			additionalTank -= 1
		}
	}

	// 最后一次剩余的燃料消耗
	totalDistance += mainTank * 10

	return totalDistance
}

// 考虑从副油箱能得到多少燃料。
// 主油箱消耗 5 升燃料可以从副油箱得到 1 升，这 1 升又加回主油箱，相当于从主油箱的初始燃料 mainTank 中每消耗 4 升，就可以从副油箱中得到一升。所以看上去可以从副油箱得到
//
//	min(mainTank/4,additionalTank)
//
// 升燃料。但这是不对的，例如 mainTank=8，只能从副油箱得到 1 升燃料。因为 8−5+1=4 ，此时无法再从副油箱中得到燃料。
// 将上式分子减一就可以修复这个问题，也就是
//
//	min(mainTank-1/4,additionalTank)
//
// 对于 mainTank=8 只能算出 1，而对于 mainTank=9，则可以恰好从副油箱得到 2 升燃料（如果有的话）。
// 所以答案为
//
//	(mainTank+min(mainTank-1/4,additionalTank))-10
func distanceTraveled2(mainTank int, additionalTank int) int {
	return (mainTank + int(math.Min(float64(mainTank-1/4), float64(additionalTank)))) * 10
}
