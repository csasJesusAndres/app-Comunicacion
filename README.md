# Voz P2P

Aplicación de escritorio Wails para llamadas directas, chat privado y pantalla
compartida mediante PeerJS/WebRTC.

La interfaz existente se mantiene en `index.html` y se incrusta dentro del
binario Go. Wails actúa como contenedor nativo; PeerJS continúa proporcionando
el descubrimiento inicial y WebRTC transporta el audio, el chat y la pantalla.

## Requisitos

- Go 1.25 o posterior.
- Wails CLI v2.13 o posterior.
- Dependencias nativas de Wails para el sistema operativo objetivo.
- Conexión a Internet para cargar PeerJS y usar el servidor de señalización
  público de PeerJS.

Instalar la CLI de Wails:

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@v2.13.0
```

## Desarrollo

```bash
wails dev
```

## Tests y compilación

```bash
go test ./...
wails build
```

El ejecutable se genera en `build/bin/voz-p2p` (la extensión depende del
sistema operativo).

## Consideraciones

- El micrófono y la pantalla siguen requiriendo permisos del WebView.
- PeerJS se carga desde `https://unpkg.com`; una futura versión puede empaquetar
  una copia local si se necesita funcionamiento sin Internet.
- No se añadió lógica de negocio Go porque esta integración conserva el flujo
  P2P actual sin reescribirlo.
