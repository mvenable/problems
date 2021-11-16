package bouncing_balls

/*
BouncingBall

A child is playing with a ball on the nth floor of a tall building. The height of this floor, h, is known.

He drops the ball out of the window. The ball bounces (for example), to two-thirds of its height (a bounce of 0.66).

His mother looks out of a window 1.5 meters from the ground.

How many times will the mother see the ball pass in front of her window (including when it's falling and bouncing?)
Three conditions must be met for a valid experiment:

    Float parameter "h" in meters must be greater than 0
    Float parameter "bounce" must be greater than 0 and less than 1
    Float parameter "window" must be less than h.

If all three conditions above are fulfilled, return a positive integer, otherwise return -1.

NOTE: The CodeWars implementation expects the total times viewed to always be incremented by 2 (one up and one down).
      I diverged from this implementation.  If the ball bounces to exactly the height of the window it will only be
      seen once.
*/
func BouncingBall(floorHeight, bounceMultiplier, windowHeight float64) int {
	if floorHeight <= 0 || bounceMultiplier >= 1 || bounceMultiplier <= 0 || windowHeight >= floorHeight {
		return -1
	}

	totalViews := 1
	currentHeight := floorHeight
	for {
		currentHeight *= bounceMultiplier
		if currentHeight > windowHeight {
			totalViews += 2
		} else if currentHeight == windowHeight {
			totalViews += 1
		} else {
			break
		}
	}

	return totalViews
}
