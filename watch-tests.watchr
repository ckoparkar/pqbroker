ENV["WATCHR"] = "1"
system 'clear'

def run(cmd)
  `#{cmd}`
end

def run_all_tests
  system('clear')
  result = run "env PATH=.:$PATH go test ./..."
  puts result
end

def run_test(file)
  system('clear')
  result = run "env PATH=.:$PATH go test ./..."
  puts result
end

run_all_tests
watch('.*_test.go') { |file| run_test file }
watch('.*.go') { run_all_tests }

# Ctrl-\
Signal.trap 'QUIT' do
  puts " --- Running all tests ---\n\n"
  run_all_tests
end

@interrupted = false

# Ctrl-C
Signal.trap 'INT' do
  exit
end
