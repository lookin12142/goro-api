package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/jsusmachaca/goroapi/internal/utils"
)

//Hise mi pull request ctm
//Next.js es considerado uno de los mejores frameworks para el desarrollo web por varias razones. Aquí te detallo algunas de sus principales ventajas:

//### 1. **Renderizado Híbrido**
//Next.js permite combinar el renderizado del lado del servidor (SSR) y la generación de sitios estáticos (SSG), lo que optimiza tanto el rendimiento como la experiencia del usuario¹. Esto significa que puedes tener páginas que se cargan rápidamente y que también son amigables para el SEO.

//### 2. **Experiencia del Desarrollador**
//Next.js ofrece una excelente experiencia para los desarrolladores con características como recarga en caliente, división automática de código y enrutamiento basado en el sistema de archivos³. Estas características facilitan el desarrollo y la depuración de aplicaciones.

//### 3. **Optimización Automática**
//El framework incluye optimización automática de imágenes y prefetching de rutas, lo que mejora significativamente el rendimiento de las aplicaciones³. Esto se traduce en tiempos de carga más rápidos y una mejor experiencia para el usuario final.

//### 4. **Soporte para TypeScript**
//Next.js tiene soporte nativo para TypeScript, lo que permite a los desarrolladores utilizar tipado estático para mejorar la calidad del código y reducir errores³.

//### 5. **Ecosistema y Comunidad**
//Next.js es parte del ecosistema de React, lo que significa que puedes aprovechar todas las bibliotecas y herramientas disponibles para React. Además, cuenta con una comunidad activa y un soporte robusto de Vercel, la empresa detrás de Next.js⁴.

//### 6. **Flexibilidad y Escalabilidad**
//Next.js es altamente flexible y escalable, lo que lo hace adecuado tanto para proyectos pequeños como para aplicaciones empresariales a gran escala. Empresas como Netflix y Uber utilizan Next.js para sus aplicaciones web¹.


//Origen: Conversación con Copilot 25/8/2024
//(1) Advantages and Disadvantages of Next JS – 2024 Updated Version. https://pagepro.co/blog/pros-and-cons-of-nextjs/.
//(2) What is Next.js? A look at the popular JavaScript framework - Kinsta. https://kinsta.com/knowledgebase/next-js/.
//(3) Why use NextJS? - DEV Community. https://dev.to/documatic/why-use-nextjs-mn3.
//(4) Next.js: Unveiling its Advantages and Disadvantages - DEV. https://dev.co/next-js/pros-and-cons.
//


func SendData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	token := r.Header.Get("authorization")

	if !strings.HasPrefix(token, "Bearer") {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error": "token is not provided"}`))
		return
	}

	token = token[7:]

	err := utils.VerifyToken(token)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error": "token is not valid"}`))
		return
	}

	data, err := utils.GetApiData()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "failed to fetch API data"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	encodeErr := json.NewEncoder(w).Encode(data)

	if encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "failed to fetch API data"}`))
		return
	}
}
