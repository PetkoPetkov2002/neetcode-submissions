class Solution:
    def floodFill(self, image: List[List[int]], sr: int, sc: int, color: int) -> List[List[int]]:
        colourOld = image[sr][sc]
        memory = {}
        return self.recurciveFloodFill(image, sr, sc, color, memory, colourOld)
    
    def recurciveFloodFill(self, image: List[List[int]], r: int, c: int, colourNew: int, memory: dict, colourOld: int) -> List[List[int]]:
        if r >= len(image) or c >= len(image[0]):
            return image
        if r < 0 or c < 0:
            return image
        if image[r][c] != colourOld:
            return image
        if (r, c) in memory:
            return image
        
        image[r][c] = colourNew
        memory[(r, c)] = True
        
        self.recurciveFloodFill(image, r+1, c, colourNew, memory, colourOld)
        self.recurciveFloodFill(image, r-1, c, colourNew, memory, colourOld)
        self.recurciveFloodFill(image, r, c+1, colourNew, memory, colourOld)
        self.recurciveFloodFill(image, r, c-1, colourNew, memory, colourOld)
            
        return image