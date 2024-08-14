package face

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"os"
	"strings"
)

var (
	red    = lipgloss.AdaptiveColor{Light: "#FE5F86", Dark: "#FE5F86"}
	indigo = lipgloss.AdaptiveColor{Light: "#5A56E0", Dark: "#7571F9"}
	green  = lipgloss.AdaptiveColor{Light: "#02BA84", Dark: "#02BF87"}
)

type Styles struct {
	Base,
	HeaderText,
	Status,
	StatusHeader lipgloss.Style
}

func NewStyles(lg *lipgloss.Renderer) *Styles {
	s := Styles{}
	s.Base = lg.NewStyle().
		Padding(1, 4, 0, 1)
	s.HeaderText = lg.NewStyle().
		Foreground(indigo).
		Bold(true).
		Padding(0, 1, 0, 2)
	s.Status = lg.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(indigo).
		PaddingLeft(1).
		MarginTop(1)
	s.StatusHeader = lg.NewStyle().
		Foreground(green).
		Bold(true)
	return &s
}

type state int

const (
	statusNormal state = iota
	stateDone
)

type formModel struct {
	form   *huh.Form
	lg     *lipgloss.Renderer
	styles *Styles
	width  int
}

func newFormModel(filenames []string) formModel {
	// if opts.mode == "executar" {
	// 	if opts.route == "cliente_contrato" {
	// 		sel := &survey.Select{Message: "Qual o campo?", Options: formModel.GetContratosFields()}
	// 		error = survey.AskOne(sel, &opts.field)
	// 		exit_if_error(error, sel.Message)
	// 		prom := &survey.Input{Message: "Qual o valor?"}
	// 		error = survey.AskOne(prom, &opts.value)
	// 		exit_if_error(error, prom.Message)
	// 	} else {
	// 		fmt.Print("Rota não implementada!")
	// 	}
	// }

	modeSel := huh.NewSelect[string]()
	modeSel.Title("O que fazer?")
	modeSel.Options(huh.NewOptions("visualização", "execução")...)
	modeSel.Key("mode")

	routeSel := huh.NewSelect[string]()
	routeSel.Title("Em qual a rota?")
	routeSel.Options(huh.NewOptions("cliente", "cliente_contrato")...)
	routeSel.Key("rota")

	envSel := huh.NewSelect[string]()
	envSel.Title("Em qual o ambiente?")
	envSel.Options(huh.NewOptions("desenvolvimento", "produção")...)
	envSel.Key("env")

	descInput := huh.NewText()
	descInput.Title("Adicione uma descrição!")
	descInput.Key("desc")

	fileSel := huh.NewSelect[string]()
	fileSel.Title("Qual o arquivo 'csv'?")
	fileSel.Options(huh.NewOptions(filenames...)...)
	fileSel.Key("filename")

	form := huh.NewForm(
		huh.NewGroup(
			modeSel,
			routeSel,
			envSel,
			descInput,
			fileSel,
		),
	)

	lg := lipgloss.DefaultRenderer()
	styles := NewStyles(lg)

	return formModel{
		form:   form,
		lg:     lg,
		styles: styles,
		width:  80,
	}
}

func (m formModel) Init() tea.Cmd {
	return m.form.Init()
}

func (m formModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	var cmds []tea.Cmd

	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m formModel) View() string {
	if m.form.State == huh.StateCompleted {
		env := m.form.GetString("env")
		mode := m.form.GetString("mode")
		return fmt.Sprintf("You selected: %s and %s", env, mode)
	}

	// return m.form.View()

	v := strings.TrimSuffix(m.form.View(), "\n\n")
	form := m.lg.NewStyle().Margin(1, 0).Render(v)

	const statusWidth = 28
	statusMarginLeft := m.width - statusWidth - lipgloss.Width(form) - m.styles.Status.GetMarginRight()
	status := m.styles.Status.
		Height(lipgloss.Height(form)).
		Width(statusWidth).
		MarginLeft(statusMarginLeft).
		Render(m.styles.StatusHeader.Render("Current Build") + "\n")

	header := lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Left,
		m.styles.HeaderText.Render("A Orchestra"),
		lipgloss.WithWhitespaceChars("="),
		lipgloss.WithWhitespaceForeground(red),
	)

	body := lipgloss.JoinHorizontal(lipgloss.Top, form, status)

	return m.styles.Base.Render(header + "\n" + body + "\n\n")
}

func Menu(filenames []string) {
	_, err := tea.NewProgram(newFormModel(filenames), tea.WithAltScreen()).Run()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}
}
