package rubik

// path              current search path (acts like a stack)
// node              current node (last node in current path)
// g                 the cost to reach current node
// f                 estimated cost of the cheapest path (root..node..goal)
// h(node)           estimated cost of the cheapest path (node..goal)
// cost(node, succ)  step cost function
// is_goal(node)     goal test
// successors(node)  node expanding function, expand nodes ordered by g + h(node)
// ida_star(root)    return either NOT_FOUND or a pair with the best path and its cost

// procedure ida_star(root)
//   bound := h(root)
//   path := [root]
//   loop
// 	t := search(path, 0, bound)
// 	if t = FOUND then return (path, bound)
// 	if t = ∞ then return NOT_FOUND
// 	bound := t
//   end loop
// end procedure

// function search(path, g, bound)
//   node := path.last
//   f := g + h(node)
//   if f > bound then return f
//   if is_goal(node) then return FOUND
//   min := ∞
//   for succ in successors(node) do
// 	if succ not in path then
// 	  path.push(succ)
// 	  t := search(path, g + cost(node, succ), bound)
// 	  if t = FOUND then return FOUND
// 	  if t < min then min := t
// 	  path.pop()
// 	end if
//   end for
//   return min
// end function