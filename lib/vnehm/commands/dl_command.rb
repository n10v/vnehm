module Vnehm
  class DlCommand < Command

    def initialize
      super

      add_option(:to, 'to ПУТЬ',
                 'Скачать аудиозапись(и) в ПУТЬ')
    end

    def execute
      @options[:dl] = 'yes'

      get_cmd = CommandManager.command_instance('get')
      get_cmd.options = @options
      get_cmd.execute
    end

    def arguments
      { 'ЧИСЛО' => '(Необязательно) Скачать последние ЧИСЛО Ваших аудиозаписей' }
    end

    def program_name
      'vnehm dl'
    end

    def summary
      'Загрузка Ваших аудиозаписей из VK'
    end

    def usage
      "#{program_name} [ЧИСЛО] [ОПЦИИ]"
    end

  end
end
