package utils_test

import (
	"tabelf/backend/service/api/pkg/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"tabelf/backend/service/api/pkg/utils"
)

func TestValidWeekPeriod(t *testing.T) {
	period := [7]bool{true, true, true, true, true, true, true}
	assert.Equal(t, "周一至周日", utils.ValidWeekPeriod(period))

	period = [7]bool{false, true, true, true, true, true, true}
	assert.Equal(t, "周二至周日", utils.ValidWeekPeriod(period))

	period = [7]bool{false, true, true, true, true, true, false}
	assert.Equal(t, "周二至周六", utils.ValidWeekPeriod(period))

	period = [7]bool{false, true, true, false, true, true, false}
	assert.Equal(t, "周二至周三、周五至周六", utils.ValidWeekPeriod(period))

	period = [7]bool{false, true, false, false, true, true, false}
	assert.Equal(t, "周二、周五至周六", utils.ValidWeekPeriod(period))

	period = [7]bool{false, true, false, true, false, true, false}
	assert.Equal(t, "周二、周四、周六", utils.ValidWeekPeriod(period))

	period = [7]bool{false, false, false, false, false, false, false}
	assert.Equal(t, "", utils.ValidWeekPeriod(period))
}
