// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package comps

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"ago/cfg"
	"strconv"
)

func HeaderTpl() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<header><h1>HTMX Terra</h1>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = ConfigOptions(cfg.IslandsConfig()).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</header>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func ConfigActions() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section class=\"config-actions\"><button class=\"button -dynamic\" hx-get=\"/display\" hx-target=\"#mini-map\" hx-swap=\"outerHTML\">Get Last</button> <button class=\"button -dynamic\" hx-post=\"/new\" hx-target=\"#mini-map\" hx-swap=\"none\" hx-trigger=\"click\" hx-include=\"#config-form\" hx-on:htmx:after-request=\"window.appState.update({viewMode: &#39;2d&#39;})\">New Map</button></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func ConfigOptions(data cfg.MapConfig) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"scrollbar-clipper\"><form id=\"config-form\" class=\"config-form\"><fieldset><legend><h3>General</h3></legend><div class=\"form-field\"><label for=\"size-form\">Map Size</label><fieldset id=\"size-form\" class=\"size-form grouped-radio\" _=\"on change remove .active-label from my children\n                    then add .active-label to target.nextElementSibling\"><input class=\"hidden-radio\" type=\"radio\" id=\"small-select\" name=\"size\" value=\"s\"> <label for=\"small-select\">s</label> <input class=\"hidden-radio\" type=\"radio\" id=\"medium-select\" name=\"size\" value=\"m\"> <label for=\"medium-select\">m</label> <input class=\"hidden-radio\" type=\"radio\" id=\"large-select\" name=\"size\" value=\"l\" checked> <label for=\"large-select\" class=\"active-label\">l</label> <input class=\"hidden-radio\" type=\"radio\" id=\"huge-select\" name=\"size\" value=\"h\"> <label for=\"huge-select\">h</label></fieldset></div><div class=\"form-field\"><label for=\"PostSmoothDistance\">Smooth Distance</label> <input type=\"number\" id=\"PostSmoothDistance\" name=\"PostSmoothDistance\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(data.PostSmoothDistance))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `comps/header.templ`, Line: 62, Col: 121}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" min=\"1\" max=\"5\" step=\"1\"></div><div class=\"form-field\"><label for=\"InitialAltitude\">Base Map Height</label> <select name=\"InitialAltitude\" id=\"InitialAltitude\"><option value=\"4\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if data.InitialAltitude == 4 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" selected")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(">Mountain</option> <option value=\"2\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if data.InitialAltitude == 2 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" selected")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(">Hills</option> <option value=\"0\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if data.InitialAltitude == 0 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" selected")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(">Plains</option> <option value=\"-2\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if data.InitialAltitude == -2 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" selected")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(">Water</option> <option value=\"-4\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if data.InitialAltitude == -4 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" selected")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(">Deep Water</option></select></div></fieldset><fieldset><legend><h3>Mountains</h3></legend><div class=\"form-field\"><label for=\"Mountains\">Quantity</label> <input type=\"number\" id=\"Mountains\" name=\"Mountains\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(data.Mountains))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `comps/header.templ`, Line: 106, Col: 94}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" max=\"20\" min=\"0\"></div><div class=\"form-field\"><label for=\"MountainAltitude\">Mountain Altitude</label> <input type=\"number\" id=\"MountainAltitude\" name=\"MountainAltitude\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(data.MountainAltitude))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `comps/header.templ`, Line: 114, Col: 49}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" max=\"10\" min=\"1\"></div><div class=\"form-field\"><label for=\"MountainAltitudeWindow\">Mountain Altitude Variance</label> <input type=\"number\" id=\"MountainAltitudeWindow\" name=\"MountainAltitudeWindow\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(data.MountainAltitudeWindow))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `comps/header.templ`, Line: 125, Col: 55}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" min=\"0\" max=\"9\"></div><div class=\"form-field\"><label for=\"MountainRadius\">Mountain Radius</label> <input type=\"number\" id=\"MountainRadius\" name=\"MountainRadius\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var8 string
		templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(data.MountainRadius))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `comps/header.templ`, Line: 136, Col: 47}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" min=\"1\" max=\"10\"></div><div class=\"form-field\"><label for=\"MountainRadiusWindow\">Mountain Radius Variance</label> <input type=\"number\" id=\"MountainRadiusWindow\" name=\"MountainRadiusWindow\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var9 string
		templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(data.MountainRadiusWindow))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `comps/header.templ`, Line: 147, Col: 53}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" min=\"0\" max=\"10\"></div></fieldset><fieldset><legend><h3>Mountain Ranges/Clusters</h3></legend><div class=\"form-field\"><label for=\"MountainRanges\">Quantity</label> <input type=\"number\" id=\"MountainRanges\" name=\"MountainRanges\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var10 string
		templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(data.MountainRanges))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `comps/header.templ`, Line: 163, Col: 47}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" min=\"0\" max=\"10\"></div><div class=\"form-field\"><label for=\"MountainRangeSize\">Mountains per Cluster</label> <input type=\"number\" id=\"MountainRangeSize\" name=\"MountainRangeSize\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var11 string
		templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(data.MountainRangeSize))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `comps/header.templ`, Line: 174, Col: 50}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" min=\"1\" max=\"10\"></div><div class=\"form-field\"><label for=\"RangeSpread\">Cluster Spread</label> <input type=\"number\" id=\"RangeSpread\" name=\"RangeSpread\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var12 string
		templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(data.RangeSpread))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `comps/header.templ`, Line: 185, Col: 44}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" min=\"1\" max=\"30\"></div></fieldset><fieldset><legend><h3>Mountain Crests</h3></legend><div class=\"form-field\"><label for=\"DefaultRunners\">Quantity</label> <input type=\"number\" id=\"DefaultRunners\" name=\"DefaultRunners\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var13 string
		templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(data.DefaultRunners))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `comps/header.templ`, Line: 201, Col: 47}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" min=\"0\" max=\"20\"></div><div class=\"form-field\"><label for=\"DefaultRunnerMinlength\">Crest Min Length</label> <input type=\"number\" id=\"DefaultRunnerMinlength\" name=\"DefaultRunnerMinlength\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var14 string
		templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(data.DefaultRunnerMinlength))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `comps/header.templ`, Line: 212, Col: 55}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" min=\"0\" max=\"5\"></div><div class=\"form-field\"><label for=\"DefaultRunnerMaxlength\">Crest Max Length</label> <input type=\"number\" id=\"DefaultRunnerMaxlength\" name=\"DefaultRunnerMaxlength\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var15 string
		templ_7745c5c3_Var15, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(data.DefaultRunnerMaxlength))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `comps/header.templ`, Line: 223, Col: 55}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var15))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" min=\"0\" max=\"10\"></div></fieldset></form></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = ConfigActions().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
