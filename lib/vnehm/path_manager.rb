module Vnehm

  ##
  # Path manager works with download paths

  module PathManager

    def self.default_dl_path
      Cfg[:dl_path]
    end

    ##
    # Checks path for validation and returns it if valid

    def self.get_path(path)
      # Check path for existence
      UI.term 'Такой папки не существует! Пожалуйста, введите корректный путь' unless Dir.exist?(path)

      File.expand_path(path)
    end

    def self.set_dl_path
      loop do
        ask_sentence = 'Введите путь в желаемую директорию скачиваемых аудиозаписей'
        default_path = File.join(ENV['HOME'], '/Music')

        if Dir.exist?(default_path)
          ask_sentence << " (нажмите Enter, чтобы установить #{default_path.magenta} в качестве этой директории)"
        else
          default_path = nil
        end

        path = UI.ask(ask_sentence + ':')

        # If user press enter, set path to default
        path = default_path if path == '' && default_path

        if Dir.exist?(path)
          Cfg[:dl_path] = File.expand_path(path)
          UI.say "#{'Директория загружаемых аудиотреков установлена по пути'.green} #{path.magenta}"
          break
        else
          UI.error 'Такой папки не существует! Пожалуйста, введите корректный путь'
        end
      end
    end

  end
end