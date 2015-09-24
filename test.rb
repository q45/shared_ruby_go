require 'ffi'
require 'json'

module Scatter
	def self.request(uris)
		ScatterBinding.scatter_request(JSON.dump({uris: uris}))
	end

	module ScatterBinding
		extend FFI::Library
		ffi_lib File.expand_path("./scraper.so", File.dirname(__FILE__))
		attach_function :scatter_request, [:string], :string
	end
end


urls = %w{http://ruby-lang.org http://rubygems.org http://golang.org http://espn.go.com}

Scatter.request(urls)