import System.Environment
import Data.Char (toLower)
import Data.List ((\\))

containsAll :: String -> String
containsAll = (['a' .. 'z'] \\)
main = do
  (toTest:args) <- getArgs
  print $ containsAll toTest
