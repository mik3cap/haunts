--female cultist
function Think()
	target = pursue()
	if target == nil then
		target = retaliate()
	end
	if target == nil then
		target = targetAllyAttacker()
	end
	if target == nil then
		target = targetAllyTarget()
	end
	if target == nil then
		target = targetLowestStat("HpCur")
	end
	if target == nil then
		target = nearest()
	end	
	if target == nil then
		return
	end
	if target.Conditions["Blindness"] then
		moveWithinRangeAndAttack (1, "Envenomed Blade", target)
	else
		moveWithinRangeAndAttack (1, "Parasitic Gift", target)
	end
end
