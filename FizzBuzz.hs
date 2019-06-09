-- FizzBuss
fizzbuzz :: Int -> String
fizzbuzz n
    | n `mod` 15 == 0 = "FizzBuzz"
    | n `mod` 3 == 0 = "Fizz"
    | n `mod` 5 == 0 = "Buss"
    | otherwise = show n