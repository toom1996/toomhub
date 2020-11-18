// @Description 标签util包
// @Author: toom1996 <1023150697@qq.com>
// @dateTime: 2020/11/18 11:51
package util

func GetLevelTag(experience int) [3]string {
	var responseSlice = [3]string{}
	switch {
	//LV1
	case experience >= 0 && experience <= 5:
		return [3]string{"LV.1", "#70D8B3", "#ffffff"}
	//LV2
	case experience > 5 && experience <= 20:
		return [3]string{"LV.2", "#70D8B3", "#ffffff"}
	//LV3
	case experience > 20 && experience <= 50:
		return [3]string{"LV.3", "#70D8B3", "#ffffff"}
	//LV4
	case experience > 50 && experience <= 100:
		return [3]string{"LV.4", "#76A9E9", "#ffffff"}
	//LV5
	case experience > 100 && experience <= 200:
		return [3]string{"LV.5", "#76A9E9", "#ffffff"}
	//LV6
	case experience > 200 && experience <= 400:
		return [3]string{"LV.6", "#76A9E9", "#ffffff"}
	//LV7
	case experience > 400 && experience <= 800:
		return [3]string{"LV.7", "#76A9E9", "#ffffff"}
	//LV8
	case experience > 800 && experience <= 1600:
		return [3]string{"LV.8", "#76A9E9", "#ffffff"}
	//LV9
	case experience > 1600 && experience <= 3200:
		return [3]string{"LV.9", "#76A9E9", "#ffffff"}
	//LV10
	case experience > 3200 && experience <= 6400:
		return [3]string{"LV.10", "#FDD000", "#ffffff"}
	//LV11
	case experience > 6400 && experience <= 12800:
		return [3]string{"LV.11", "#FDD000", "#ffffff"}
	//LV12
	case experience > 12800 && experience <= 25600:
		return [3]string{"LV.12", "#FDD000", "#ffffff"}
	//LV13
	case experience > 25600 && experience <= 51200:
		return [3]string{"LV.13", "#DC143C", "#ffffff"}
	//LV14
	case experience > 51200 && experience <= 102400:
		return [3]string{"LV.14", "#DC143C", "#ffffff"}
	//LV15
	case experience > 102400 && experience <= 204800:
		return [3]string{"LV.15", "#DC143C", "#ffffff"}
	//>LV15
	case experience > 204800:
		return [3]string{"LV.MAX", "#000000", "#ffffff"}

	}
	return responseSlice
}
