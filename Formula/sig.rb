class Sig < Formula
  desc "A Simple CLI"
  homepage "https://github.com/X140Yu/homebrew-foo/blob/master/foo"
  url "", :using => :git

  def install
    # 安装源代码
    libexec.install Dir["bin/*"]
    # 创建 foo 命令
    bin.write_exec_script (libexec/"kws_sig")

    foo_path = (bin/"foo")

    # 修改 foo 的权限，让它可以被更改
    FileUtils.chmod 0755, foo_path

  end
end