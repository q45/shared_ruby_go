require 'ffi'
require 'json'
require 'scatter'
require 'benchmark'

module Scatter
	def self.request(uris)
		ScatterBinding.scatter_request(JSON.dump({uris: uris}))
	end

	module ScatterBinding
		extend FFI::Library
		ffi_lib File.expand_path("./libscatter.so", File.dirname(__FILE__))
		attach_function :scatter_request, [:string], :string
	end
end



# module Sum
# 	extend FFI::Library
# 	ffi_lib './libsum.so'
# 	# ffi_lib '/Users/q/go/src/github.com/q45/libsum/libsum.so'
# 	attach_function :add, [:int, :int], :int
# 	attach_function :fibonacci, [:int], :int
# end

urls = %w{http://ruby-lang.org http://rubygems.org http://golang.org}

# puts Sum.add(15, 27)

# puts Sum.fibonacci(10)

puts Scatter.request(urls)

# Benchmark.bmbm do |r|
# 	r.report("3 requests") do 
# 		Scatter.request(urls.shuffle.take(3))
# 	end
# end

