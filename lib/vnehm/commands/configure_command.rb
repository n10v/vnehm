module Vnehm
  class ConfigureCommand < Command

    def execute
      loop do
        show_info
        UI.newline
        show_menu
        UI.sleep
      end
    end

    def program_name
      'vnehm configure'
    end

    def summary
      'Настройка приложения'
    end

    def usage
      'vnehm configure'
    end

    private

    def show_info
      dl_path = PathManager.default_dl_path
      UI.say "Директория скачанных аудиозаписей: #{dl_path.magenta}" if dl_path

      if OS.mac?
        playlist = PlaylistManager.default_playlist
        UI.say "Плейлист iTunes: #{playlist.to_s.cyan}" if playlist
      end
    end

    def show_menu
      UI.menu do |menu|
        menu.choice(:inc, 'Изменить путь для скачиваемых аудиозаписей'.freeze) { PathManager.set_dl_path }
        menu.choice(:inc, 'Изменить плейлист iTunes'.freeze) { PlaylistManager.set_playlist } if OS.mac?
        menu.choice(:inc, 'Авторизация'.freeze) { TokenManager.authorize }
      end
    end

  end
end
