# coding: utf-8
lib = File.expand_path('../lib', __FILE__)
$LOAD_PATH.unshift(lib) unless $LOAD_PATH.include?(lib)
require 'vnehm/version'

Gem::Specification.new do |spec|
  spec.name    = 'vnehm'
  spec.version = Vnehm::VERSION
  spec.authors = ['Albert Nigmatzianov']
  spec.email   = ['albertnigma@gmail.com']

  spec.summary     = 'Convenient way to download music (also to iTunes) from VKontakte via terminal'
  spec.description = 'vnehm is a console tool, which downloads, sets IDv3 tags and adds to your iTunes library your VKontakte music in convenient way. See homepage for instructions'
  spec.homepage    = 'http://www.github.com/bogem/vnehm'
  spec.license     = 'MIT'

  spec.files                 = `git ls-files -z`.split("\x0").reject { |f| f.match(%r{^(Rakefile)/}) }
  spec.bindir                = 'bin'
  spec.executables           = 'vnehm'
  spec.require_paths         = ['lib']
  spec.required_ruby_version = '>= 1.9.3'
  spec.post_install_message  = "Don't forget to install Taglib from OS' package manager!"

  spec.add_dependency 'certifi'
  spec.add_dependency 'colored'
  spec.add_dependency 'taglib-ruby', '>= 0.7.0'
  spec.add_dependency 'vkontakte_api', '~> 1.4'
end
