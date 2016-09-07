import System.Environment
import Data.List ((\\))

findMissing :: String -> String
findMissing = (['a' .. 'z'] \\)
main = do
  (toTest:args) <- getArgs
  print $ findMissing toTest
