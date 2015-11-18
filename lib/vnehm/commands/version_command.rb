require 'vnehm/version'

module Vnehm
  class VersionCommand < Command

    def execute
      UI.say VERSION
    end

    def program_name
      'vnehm version'
    end

    def summary
      'Вывод версии приложения'
    end

    def usage
      program_name
    end

  end
end
