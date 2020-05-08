class Sig < Formula
  desc "A Simple CLI"
  homepage "https://github.com/melotusme/sig"
  # url "git@github.com:melotusme/sig.git", :using => :git
  url "https://github.com/melotusme/sig.git", :using => :git
  version "0.0.1"

  def install
    # 安装源代码
    bin.install Dir["bin/*"]
  end
end