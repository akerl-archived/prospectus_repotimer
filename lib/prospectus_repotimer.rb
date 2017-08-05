module ProspectusCircleci
  ##
  # Helper for automatically adding build status check
  class Build < Module
    def initialize(days_ago)
      @days_ago = days_ago || raise('No age specified')
    end

    def extended(other) # rubocop:disable Metrics/MethodLength
      actual_val, expected_val = parse_status

      other.deps do
        item do
          name 'timer'

          expected do
          end

          actual do
          end
        end
      end
    end
  end
end
