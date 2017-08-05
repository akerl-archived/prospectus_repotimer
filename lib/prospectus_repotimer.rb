require 'date'

module ProspectusRepotimer
  ##
  # Helper for automatically adding timer check
  class Timer < Module
    def initialize(days_ago)
      @days_ago = days_ago || raise('No age specified')
    end

    def extended(other) # rubocop:disable Metrics/MethodLength
      actual_val, expected_val = parse_status

      other.deps do
        item do
          name 'timer'

          expected do
            static
            set expected_val
          end

          actual do
            static
            set actual_val
          end
        end
      end
    end

    def parse_status
      [last_commit_days_ago, days_ago_array]
    end

    def last_commit_days_ago
      @last_commit_days_ago ||= (Date.today - last_commit).to_i
    end

    def last_commit
      @last_commit ||= Date.parse(`git show --format=%cI -s`)
    end

    def days_ago_array
      @days_ago.downto(0).to_a
    end
  end
end
