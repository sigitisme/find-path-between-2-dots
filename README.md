# Find Path Between 2 Dots
This is a solution to find path between 2 dots in a 2D array. This 2D array may have obstacles on it marked by -1 value.  

This solution is using Bread First Search algorithm. The step is traverse the array and mark the route after it is visited.

## Test Cases

### Test Case 1: 5x5 array with no obstacles
```
{0, 0, 0, 0, 0}
{0, 0, 0, 0, 0}
{0, 0, 0, 0, 0}
{0, 0, 0, 0, 0}
{0, 0, 0, 0, 0}
```
### Test Case 2:  5x5 array with little resistances
```
{0, 0, 0, 0, 0}
{0, 0, 0, 0, 0}
{0, 0, 0, 0, 0}
{0, 0, -1, 0, 0}
{0, -1, 0, 0, 0}
```
### Test Case 3: 5x5 array with high resistances
```
{0, 0, 0, 0, 0}
{-1, -1, -1, -1, 0}
{0, 0, 0, 0, 0}
{0, -1, -1, -1, -1}
{0, 0, 0, 0, 0}
```
### Test Case 4: 5x5 array with case example resistances
```
{0, 0, 0, -1, 0}
{-1, 0, 0, -1, -1}
{0, 0, 0, -1, 0}
{-1, 0, 0, 0, 0}
{0, 0, -1, 0, 0}
```
### Test Case 5: 5x4 array with case example resistances
```
{0, 0, 0, -1}
{-1, 0, 0, -1}
{0, 0, 0, -1}
{-1, 0, 0, 0}
{0, 0, -1, 0}
```
### Test Case 6: 5x5 array with no path 1
```
{0, 0, 0, 0, 0},
{-1, -1, -1, -1, -1},
{0, 0, 0, 0, 0},
{0, -1, -1, -1, -1},
{0, 0, 0, 0, 0},
```
### Test Case 7: 5x5 array with no path 2
```
{0, 0, 0, 0, 0}
{-1, -1, -1, -1, 0}
{0, 0, 0, 0, 0}
{0, -1, -1, -1, -1}
{0, 0, 0, -1, 0}
```
## Run
### run the tests
```
make test
```
### view the coverage in browser

```
make coverage
```

End of file