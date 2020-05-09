class Smex < Formula
	desc "a blazing fast cli app to extract/convert/process sitemaps"
	homepage "https://github.com/hbish/smex"
	version "v0.0.4"

	if OS.mac?
		url "https://github.com/hbish/smex/releases/download/v0.0.4/smex_v0.0.4_darwin_amd64.tar.gz"
		sha256 "5ec9fe1e6bd92e6522f9ccd3b2f0728398b75b0fdc8813d648907f04c8beb06f"
	elsif OS.linux?
		if Hardware::CPU.is_64_bit?
			url "https://github.com/hbish/smex/releases/download/v0.0.4/smex_v0.0.4_linux_amd64.tar.gz"
			sha256 "f6a40f3ec29ae9e69261eedb2518a056633a5ae7690816efba55ba7f0fd5f249"
		else
			url "https://github.com/hbish/smex/releases/download/v0.0.4/smex_v0.0.4_linux_386.tar.gz"
			sha256 "344ca898fb3e55726145937eb378f30d6a7dd677ab21e13689ee137625ceabda"
		end
	end

	def install
		bin.install "smex"
	end
end

