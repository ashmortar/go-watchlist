// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.432
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Contact() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"mx-auto sm:p-0 my-auto main-card\"><section class=\"mx-auto bg-white max-w-screen-xl rounded-lg shadow-md\"><div class=\"grid grid-cols-4 text-gray-800\"><div class=\"order-first col-span-4 max-w-screen-md px-8 py-10 md:order-last md:px-10 md:py-12\"><h2 class=\"mb-8 text-2xl font-black\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var2 := `Get in touch`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var2)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h2><p class=\"mt-2 mb-4 font-sans text-sm tracking-normal\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var3 := `Don't be shy to ask me a question.`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var3)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><form action=\"\"><div class=\"mb-5 grid md:grid-cols-11\"><input class=\"mb-2 md:mb-0 col-span-2 md:col-span-5 bg-orange-100 shadow-md w-full border-b py-3 text-sm outline-none focus:border-b-2 focus:border-amber-500\" type=\"text\" placeholder=\"Name\" name=\"name\"><div class=\"hidden md:inline col-span-1\"></div><input class=\"col-span-2 md:col-span-5 bg-orange-100 shadow-md w-full border-b py-3 text-sm outline-none focus:border-b-2 focus:border-amber-500\" type=\"email\" placeholder=\"Email\" name=\"email\"></div><textarea class=\"mb-10 w-full bg-orange-100 shadow-md resize-y whitespace-pre-wrap border-b py-3 text-sm outline-none focus:border-b-2 focus:border-amber-500\" id=\"\" rows=\"6\" placeholder=\"Question\" name=\"question\"></textarea> <button type=\"submit\" class=\"group flex cursor-pointer items-center rounded-xl bg-teal-500 bg-none px-8 py-4 text-center font-semibold leading-tight text-white\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var4 := `Send`
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var4)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" <svg class=\"group-hover:ml-8 ml-4 transition-all\" xmlns=\"http://www.w3.org/2000/svg\" aria-hidden=\"true\" role=\"img\" width=\"1em\" height=\"1em\" preserveAspectRatio=\"xMidYMid meet\" viewBox=\"0 0 24 24\"><path fill=\"none\" stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M9.912 12H4L2.023 4.135A.662.662 0 0 1 2 3.995c-.022-.721.772-1.221 1.46-.891L22 12L3.46 20.896c-.68.327-1.464-.159-1.46-.867a.66.66 0 0 1 .033-.186L3.5 15\"></path></svg></button></form></div></div></section></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
