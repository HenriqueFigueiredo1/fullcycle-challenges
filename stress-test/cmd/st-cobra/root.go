package st

/*
O sistema deverá gerar um relatório com informações específicas após a execução dos testes.

Entrada de Parâmetros via CLI:

--url: URL do serviço a ser testado.
--requests: Número total de requests.
--concurrency: Número de chamadas simultâneas.
*/

import (
	"fmt"
	"net/url"
	"os"

	"github.com/HenriqueFigueiredo1/goexpert-desafio-stress-test/internal/entity"
	"github.com/HenriqueFigueiredo1/goexpert-desafio-stress-test/internal/usecase"
	"github.com/spf13/cobra"
)

var stressTestConfig *entity.StressTestConfig

var rootCmd = &cobra.Command{
	Use:   "go-stress-test",
	Short: "A stress tester CLI like Apache ab",
	Long:  `A stress tester CLI like Apache ab`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		flagUrl, _ := cmd.Flags().GetString("url")

		if flagUrl == "" {
			return fmt.Errorf("flag url is required")
		}
		_, err := url.ParseRequestURI(flagUrl)
		if err != nil {
			return fmt.Errorf("url flag is invalid")
		}

		flagConcurrency, _ := cmd.Flags().GetInt("concurrency")
		flagRequests, _ := cmd.Flags().GetInt("requests")

		if flagRequests <= 0 {
			return fmt.Errorf("requests must be greater or equals 1")
		}

		if flagConcurrency <= 0 {
			return fmt.Errorf("concurrency must be greater or equals 1")
		}

		if flagRequests < flagConcurrency {
			return fmt.Errorf("requests must be greater or equal concurrent")
		}

		stressTestConfig = entity.NewStressTestConfig(flagUrl, flagConcurrency, flagRequests)

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		st := usecase.NewExecuteStressTestUseCase()
		results := st.Execute(stressTestConfig)
		results.PrintReport()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	var flagUrl string
	var flagConcurrency, flagRequests int
	rootCmd.Flags().StringVarP(&flagUrl, "url", "u", "", "URL to stress")
	rootCmd.Flags().IntVarP(&flagConcurrency, "concurrency", "c", 1, "URL to stress")
	rootCmd.Flags().IntVarP(&flagRequests, "requests", "r", 1, "URL to stress")
}
