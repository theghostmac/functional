countdown :: Int -> IO ()
countdown 0 = putStrLn "Blast off!"
countdown n = do
  putStrLn (show n)
  countdown (n - 1)
