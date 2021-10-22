package acorn

import (
	"github.com/arskang/gomail-acorn-template/acornstyles"
	"github.com/arskang/gomail-acorn-template/acorntypes"
)

func (h HTML) NewSpacer() string {
	return `<div class="spacer py-sm-16" style="line-height: 32px;">&zwnj;</div>`
}

func (h HTML) NewDivider(color *acorntypes.Color) string {
	acornColors := acornstyles.GetColors()
	c := acornColors.Grey.M400
	if color != nil {
		c = color
	}
	return `
	<table cellpadding="0" cellspacing="0" role="presentation" width="100%">
		<tr>
			<td class="divider py-sm-24" style="padding: 16px;">
				<div style="background: ` + c.String() + `; height: 1px; line-height: 1px;" />
			</td>
		</tr>
	</table>
	`
}
